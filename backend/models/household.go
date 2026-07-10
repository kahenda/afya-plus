package models

import "time"

type Household struct {
	ID           int       `json:"id"`
	HeadName     string    `json:"head_name"`
	Location     string    `json:"location"`
	MemberCount  int       `json:"member_count"`
	CreatedBy    int       `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateHouseholdInput struct {
	HeadName    string `json:"head_name" binding:"required"`
	Location    string `json:"location"`
	MemberCount int    `json:"member_count"`
	CreatedBy   int    `json:"created_by" binding:"required"`
}
