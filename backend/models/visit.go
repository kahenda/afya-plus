package models

import "time"

type Visit struct {
	ID                  int       `json:"id"`
	HouseholdID         int       `json:"household_id"`
	ChwID               int       `json:"chw_id"`
	ClientVisitID       string    `json:"client_visit_id"`
	VisitType           string    `json:"visit_type"`
	ChildWeightKg       *float64  `json:"child_weight_kg,omitempty"`
	ChildAgeMonths      *int      `json:"child_age_months,omitempty"`
	VaccinationDueDate  *string   `json:"vaccination_due_date,omitempty"`
	VaccinationDone     bool      `json:"vaccination_done"`
	Notes               string    `json:"notes"`
	VisitedAt           time.Time `json:"visited_at"`
	SyncedAt            time.Time `json:"synced_at"`
}

type CreateVisitInput struct {
	HouseholdID        int      `json:"household_id" binding:"required"`
	ChwID              int      `json:"chw_id" binding:"required"`
	ClientVisitID      string   `json:"client_visit_id" binding:"required"`
	VisitType          string   `json:"visit_type"`
	ChildWeightKg      *float64 `json:"child_weight_kg"`
	ChildAgeMonths     *int     `json:"child_age_months"`
	VaccinationDueDate *string  `json:"vaccination_due_date"`
	VaccinationDone    bool     `json:"vaccination_done"`
	Notes              string   `json:"notes"`
	VisitedAt          string   `json:"visited_at" binding:"required"`
}
