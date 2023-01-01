package domain

import "github.com/google/uuid"

type CartProduct struct {
	ID         uuid.UUID
	ProductID  uuid.UUID
	CustomerID uuid.UUID
	Name       string
	Price      float64
	PhotoUrl   string
	Quantity   int
}
