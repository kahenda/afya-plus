package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kahenda/afya-plus/backend/config"
	"github.com/kahenda/afya-plus/backend/models"
)

// checkAndCreateFlags applies simple risk rules after a visit is logged.
// NOTE: the malnutrition threshold here is a simplified placeholder, not a
// real WHO growth-standard calculation. Refine with proper weight-for-age
// percentile tables before relying on this for real clinical decisions.
func checkAndCreateFlags(visitID int, householdID int, input models.CreateVisitInput) {
	ctx := context.Background()

	// Rule 1: possible malnutrition risk
	if input.ChildWeightKg != nil && input.ChildAgeMonths != nil {
		expectedMinWeight := float64(*input.ChildAgeMonths)*0.2 + 2.5
		if *input.ChildWeightKg < expectedMinWeight {
			insertFlag(ctx, visitID, householdID, "malnutrition_risk")
		}
	}

	// Rule 2: overdue vaccination
	if input.VaccinationDueDate != nil && !input.VaccinationDone {
		dueDate, err := time.Parse("2006-01-02", *input.VaccinationDueDate)
		if err == nil && dueDate.Before(time.Now()) {
			insertFlag(ctx, visitID, householdID, "overdue_vaccination")
		}
	}
}

func insertFlag(ctx context.Context, visitID int, householdID int, reason string) {
	query := `
		INSERT INTO flags (visit_id, household_id, reason)
		VALUES ($1, $2, $3)
	`
	config.DB.Exec(ctx, query, visitID, householdID, reason)
}

func GetFlags(c *gin.Context) {
	statusFilter := c.Query("status")

	var rows interface {
		Close()
		Next() bool
		Scan(dest ...interface{}) error
	}
	var err error

	ctx := context.Background()
	if statusFilter != "" {
		rows, err = config.DB.Query(ctx, `
			SELECT id, visit_id, household_id, reason, status, assigned_to, created_at, updated_at
			FROM flags WHERE status = $1 ORDER BY created_at DESC
		`, statusFilter)
	} else {
		rows, err = config.DB.Query(ctx, `
			SELECT id, visit_id, household_id, reason, status, assigned_to, created_at, updated_at
			FROM flags ORDER BY created_at DESC
		`)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch flags: " + err.Error()})
		return
	}
	defer rows.Close()

	flags := []models.Flag{}
	for rows.Next() {
		var f models.Flag
		if err := rows.Scan(&f.ID, &f.VisitID, &f.HouseholdID, &f.Reason, &f.Status, &f.AssignedTo, &f.CreatedAt, &f.UpdatedAt); err != nil {
			continue
		}
		flags = append(flags, f)
	}

	c.JSON(http.StatusOK, flags)
}

func UpdateFlagStatus(c *gin.Context) {
	flagID := c.Param("id")
	var input models.UpdateFlagStatusInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE flags SET status = $1, assigned_to = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING id, visit_id, household_id, reason, status, assigned_to, created_at, updated_at
	`

	var f models.Flag
	err := config.DB.QueryRow(context.Background(), query, input.Status, input.AssignedTo, flagID).
		Scan(&f.ID, &f.VisitID, &f.HouseholdID, &f.Reason, &f.Status, &f.AssignedTo, &f.CreatedAt, &f.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update flag: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, f)
}
