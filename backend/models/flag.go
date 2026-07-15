package models

import "time"

type Flag struct {
	ID          int       `json:"id"`
	VisitID     int       `json:"visit_id"`
	HouseholdID int       `json:"household_id"`
	Reason      string    `json:"reason"`
	Status      string    `json:"status"`
	AssignedTo  *int      `json:"assigned_to,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateFlagStatusInput struct {
	Status     string `json:"status" binding:"required,oneof=flagged under_review action_taken resolved"`
	AssignedTo *int   `json:"assigned_to"`
}
