package models

import "gorm.io/gorm"

type Product struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `json:"name"`
	Price     float64        `json:"price"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
