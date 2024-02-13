package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Name     *string
	Password string
	Laporan  []Laporan
}

type Rekening struct {
	NomorRekening string    `gorm:"type:varchar(255);primaryKey"`
	Laporan       []Laporan `gorm:"foreignKey:NomorRekeningID"`
}

type PredictionText struct {
	gorm.Model
	Text string `gorm:"type:longtext" json:"text"`
}

type Laporan struct {
	ID              string   `gorm:"type:char(36);primaryKey"`
	NomorRekening   Rekening `gorm:"foreignKey:NomorRekeningID;references:NomorRekening"`
	NomorRekeningID string   `gorm:"type:varchar(255)"`
	Title           string
	Body            string
	Evidence        *string
	Reporter        User `gorm:"foreignKey:UserID;references:ID"`
	UserID          int
	Status          *string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
