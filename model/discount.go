package model

import (
	"time"
	"gorm.io/gorm"
)

type Discount struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Price       float64   `json:"price"`
	Discount    float64   `json:"discount"`
	FinalPrice  float64   `json:"final_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
