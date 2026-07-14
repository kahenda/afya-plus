package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahenda/afya-plus/backend/config"
	"github.com/kahenda/afya-plus/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	var pinHash string

	query := `
		SELECT id, name, phone_number, pin_hash, role, zone, created_at
		FROM users
		WHERE phone_number = $1
	`

	err := config.DB.QueryRow(context.Background(), query, input.PhoneNumber).
		Scan(&user.ID, &user.Name, &user.PhoneNumber, &pinHash, &user.Role, &user.Zone, &user.CreatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid phone number or PIN"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "login failed: " + err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(pinHash), []byte(input.Pin)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid phone number or PIN"})
		return
	}

	c.JSON(http.StatusOK, user)
}
