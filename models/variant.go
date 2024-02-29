package models

import "time"

type Variant struct {
	ID          uint
	VariantName string
	Quantity    int
	ProductID   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    int
}

type VariantResponse struct {
	ID          uint
	VariantName string
	Quantity    int
	ProductID   int
}
