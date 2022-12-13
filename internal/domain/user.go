package domain

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
}