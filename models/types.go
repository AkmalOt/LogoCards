package models

import (
	"time"
)

//type mediaType string
//
//const (
//	PHOTO mediaType = "photo"
//	VIDEO mediaType = "video"
//)

type Config struct {
	LocalHost struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
}

type DbData struct {
	DbConnection struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	}
}

type UserCards struct {
	ID        string    `json:"id" gorm:"column:id"`
	Num       int       `json:"num" gorm:"num"`
	Name      string    `json:"name" gorm:"name"`
	Logo      string    `json:"logo" gorm:"logo"`
	Active    bool      `json:"active" gorm:"active"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

//type Usercard struct {
//	Num       int       `json:"num" gorm:"num"`
//	Name      string    `json:"name" gorm:"name"`
//	Logo      string    `json:"logo" gorm:"logo"`
//	Active    bool      `json:"active" gorm:"active"`
//	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
//	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
//}
//
//type GetUser struct {
//	ID        string      `json:"id" gorm:"id"`
//	Num       int64       `json:"num" gorm:"num"`
//	Name      string      `json:"name" gorm:"name"`
//	Logo      image.Image `json:"logo" gorm:"logo"`
//	Active    bool        `json:"active" gorm:"active"`
//	CreatedAt time.Time   `json:"created_at" gorm:"created_at"`
//	UpdatedAt time.Time   `json:"updated_at" gorm:"updated_at"`
//}
