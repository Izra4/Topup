package entity

import "gorm.io/gorm"

type ListTopUp struct {
	gorm.Model
	GamesID    uint
	JenisPaket string `gorm:"type:VARCHAR(30);NOT NULL" json:"jenisPaket"`
	Harga      string `gorm:"type:VARCHAR(20);NOT NULL" json:"harga"`
	Link       string `gorm:"type:TEXT;NOT NULL" json:"link"`
	Bookings   []Booking
}
