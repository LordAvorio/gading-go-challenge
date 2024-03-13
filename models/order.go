package models

import (
	"time"
)

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `gorm:"type:varchar(100)" json:"customerName"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
	OrderedAt    time.Time `json:"orderedAt"`
	CreatedAt    *time.Time
	UpdateAt     *time.Time
}

type OrderRequest struct {
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items"`
	OrderedAt    time.Time `json:"orderedAt"`
}

type OrderResponse struct {
	CustomerName string         `json:"customerName"`
	Items        []ItemResponse `json:"items"`
	OrderedAt    time.Time      `json:"orderedAt"`
}
