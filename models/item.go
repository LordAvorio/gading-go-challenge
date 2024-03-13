package models

import (
	"time"
)

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100); unique" json:"name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     uint
	Order       *Order
	CreatedAt   *time.Time
	UpdateAt    *time.Time
}

type ItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ItemResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
