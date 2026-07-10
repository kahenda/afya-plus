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
