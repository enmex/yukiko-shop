package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Category    Category
	CreatedAt   time.Time
	Price       float64
	PhotoURL    *string
}
