package models

import "time"

// Building представляет информацию о строении
type Building struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	Name     string `json:"name" gorm:"not null"`
	City     string `json:"city" gorm:"not null"`
	Year     int    `json:"year" gorm:"not null"`
	Floors   int    `json:"floors" gorm:"not null"`
}