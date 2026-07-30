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

func CreateUser(c *gin.Context) {
	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(input.Pin), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to secure PIN"})
		return
	}

	var newUser models.User
	query := `
		INSERT INTO users (name, phone_number, pin_hash, role, zone)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, phone_number, role, zone, created_at
	`

	err = config.DB.QueryRow(context.Background(), query,
		input.Name, input.PhoneNumber, string(hashedPin), input.Role, input.Zone,
	).Scan(&newUser.ID, &newUser.Name, &newUser.PhoneNumber, &newUser.Role, &newUser.Zone, &newUser.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func GetUsers(c *gin.Context) {
	roleFilter := c.Query("role")
	ctx := context.Background()

	var rows pgx.Rows
	var err error

	if roleFilter != "" {
		rows, err = config.DB.Query(ctx, `
			SELECT id, name, phone_number, role, zone, created_at
			FROM users WHERE role = $1 ORDER BY created_at DESC
		`, roleFilter)
	} else {
		rows, err = config.DB.Query(ctx, `
			SELECT id, name, phone_number, role, zone, created_at
			FROM users ORDER BY created_at DESC
		`)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users: " + err.Error()})
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Role, &u.Zone, &u.CreatedAt); err != nil {
			continue
		}
		users = append(users, u)
	}

	c.JSON(http.StatusOK, users)
}
