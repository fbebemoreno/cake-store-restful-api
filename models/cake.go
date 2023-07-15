package models

import (
	"time"
)

type Cake struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Rating      float64   `json:"rating"`
	Image       string    `gorm:"type:varchar(1000)" json:"image"`
	Created_At  time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	Updated_At  time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`
	Deleted_At  time.Time `gorm:"index" json:"deleted_at"`
}
