package models

import (
	"time"
)

type Product struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  int
}

type ProductResponse struct {
	ID   uint
	Name string
}

type ProductWithVariantResponse struct{
	ProductName string
	VariantName string
	Quantity int
}
