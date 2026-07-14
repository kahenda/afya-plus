package models

import "time"

type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	PinHash     string    `json:"-"`
	Role        string    `json:"role"`
	Zone        string    `json:"zone"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateUserInput struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Pin         string `json:"pin" binding:"required,len=4"`
	Role        string `json:"role" binding:"required,oneof=chw supervisor"`
	Zone        string `json:"zone"`
}

type LoginInput struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Pin         string `json:"pin" binding:"required,len=4"`
}
