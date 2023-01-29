package repository

import (
	"time"
)

type UserCards struct {
	ID        string    `json:"id" gorm:"column:id"`
	Num       int       `json:"num" gorm:"num"`
	Name      string    `json:"name" gorm:"name"`
	Logo      string    `json:"logo" gorm:"logo"`
	Active    bool      `json:"active" gorm:"active"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}
