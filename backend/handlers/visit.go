package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahenda/afya-plus/backend/config"
	"github.com/kahenda/afya-plus/backend/models"
)

func CreateVisit(c *gin.Context) {
	var input models.CreateVisitInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newVisit models.Visit
	query := `
		INSERT INTO visits (household_id, chw_id, client_visit_id, visit_type, child_weight_kg, child_age_months, vaccination_due_date, vaccination_done, notes, visited_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (client_visit_id) DO NOTHING
		RETURNING id, household_id, chw_id, client_visit_id, visit_type, vaccination_done, notes, visited_at, synced_at
	`

	err := config.DB.QueryRow(context.Background(), query,
		input.HouseholdID, input.ChwID, input.ClientVisitID, input.VisitType,
		input.ChildWeightKg, input.ChildAgeMonths, input.VaccinationDueDate,
		input.VaccinationDone, input.Notes, input.VisitedAt,
	).Scan(&newVisit.ID, &newVisit.HouseholdID, &newVisit.ChwID, &newVisit.ClientVisitID,
		&newVisit.VisitType, &newVisit.VaccinationDone, &newVisit.Notes, &newVisit.VisitedAt, &newVisit.SyncedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			// ON CONFLICT DO NOTHING triggered — this visit was already synced before
			c.JSON(http.StatusOK, gin.H{"message": "visit already synced", "client_visit_id": input.ClientVisitID})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create visit: " + err.Error()})
		return
	}

	// Run risk-flagging rules now that the visit is safely stored
	checkAndCreateFlags(newVisit.ID, newVisit.HouseholdID, input)

	c.JSON(http.StatusCreated, newVisit)
}
