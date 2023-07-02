package entity

import "gorm.io/gorm"

type Games struct {
	gorm.Model
	Nama       string `gorm:"type: VARCHAR(50);NOT NULL" json:"nama"`
	Developer  string `gorm:"type: VARCHAR(50);NOT NULL" json:"developer"`
	Gambar     string `gorm:"type:TEXT;NOT NULL" json:"gambar"`
	Bookings   []Booking
	ListTopUps []ListTopUp
}
