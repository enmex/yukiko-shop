package domain

import "github.com/google/uuid"

type AccessType string

var (
	ADMIN    AccessType = "ADMIN"
	MANAGER  AccessType = "MANAGER"
	CUSTOMER AccessType = "CUSTOMER"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Role      AccessType `json:"access_type"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Password  string     `json:"password"`
}
