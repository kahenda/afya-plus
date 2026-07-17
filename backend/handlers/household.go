package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahenda/afya-plus/backend/config"
	"github.com/kahenda/afya-plus/backend/models"
)

func CreateHousehold(c *gin.Context) {
	var input models.CreateHouseholdInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.MemberCount == 0 {
		input.MemberCount = 1
	}

	var newHousehold models.Household
	query := `
		INSERT INTO households (head_name, location, member_count, created_by)
		VALUES ($1, $2, $3, $4)
		RETURNING id, head_name, location, member_count, created_by, created_at
	`

	err := config.DB.QueryRow(context.Background(), query,
		input.HeadName, input.Location, input.MemberCount, input.CreatedBy,
	).Scan(&newHousehold.ID, &newHousehold.HeadName, &newHousehold.Location, &newHousehold.MemberCount, &newHousehold.CreatedBy, &newHousehold.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create household: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newHousehold)
}

func GetHouseholds(c *gin.Context) {
	ctx := context.Background()
	rows, err := config.DB.Query(ctx, `
		SELECT id, head_name, location, member_count, created_by, created_at
		FROM households ORDER BY created_at DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch households: " + err.Error()})
		return
	}
	defer rows.Close()

	households := []models.Household{}
	for rows.Next() {
		var h models.Household
		if err := rows.Scan(&h.ID, &h.HeadName, &h.Location, &h.MemberCount, &h.CreatedBy, &h.CreatedAt); err != nil {
			continue
		}
		households = append(households, h)
	}

	c.JSON(http.StatusOK, households)
}

func GetHouseholdDetail(c *gin.Context) {
	householdID := c.Param("id")
	ctx := context.Background()

	var household models.Household
	err := config.DB.QueryRow(ctx, `
		SELECT id, head_name, location, member_count, created_by, created_at
		FROM households WHERE id = $1
	`, householdID).Scan(&household.ID, &household.HeadName, &household.Location, &household.MemberCount, &household.CreatedBy, &household.CreatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "household not found"})
		return
	}

	visitRows, err := config.DB.Query(ctx, `
		SELECT id, household_id, chw_id, client_visit_id, visit_type, vaccination_done, notes, visited_at, synced_at
		FROM visits WHERE household_id = $1 ORDER BY visited_at DESC
	`, householdID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch visits: " + err.Error()})
		return
	}
	defer visitRows.Close()

	visits := []models.Visit{}
	for visitRows.Next() {
		var v models.Visit
		if err := visitRows.Scan(&v.ID, &v.HouseholdID, &v.ChwID, &v.ClientVisitID, &v.VisitType, &v.VaccinationDone, &v.Notes, &v.VisitedAt, &v.SyncedAt); err != nil {
			continue
		}
		visits = append(visits, v)
	}

	flagRows, err := config.DB.Query(ctx, `
		SELECT id, visit_id, household_id, reason, status, assigned_to, created_at, updated_at
		FROM flags WHERE household_id = $1 ORDER BY created_at DESC
	`, householdID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch flags: " + err.Error()})
		return
	}
	defer flagRows.Close()

	flags := []models.Flag{}
	for flagRows.Next() {
		var f models.Flag
		if err := flagRows.Scan(&f.ID, &f.VisitID, &f.HouseholdID, &f.Reason, &f.Status, &f.AssignedTo, &f.CreatedAt, &f.UpdatedAt); err != nil {
			continue
		}
		flags = append(flags, f)
	}

	c.JSON(http.StatusOK, gin.H{
		"household": household,
		"visits":    visits,
		"flags":     flags,
	})
}
