package entity

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	GamesID       uint   `gorm:"NOT NULL"`
	ListTopUpId   uint   `gorm:"NOT NULL"`
	UserId        string `gorm:"type:VARCHAR(20);NOT NULL"`
	ServerId      string `gorm:"type:VARCHAR(20);NOT NULL"`
	PaymentMethod string `gorm:"type:VARCHAR(50);NOT NULL"`
	VirtualAcc    string `gorm:"type:VARCHAR(30);NOT NULL"`
	NoWA          string `gorm:"type:VARCHAR(20);NOT NULL"`
	NameAcc       string `gorm:"type:VARCHAR(20);NOT NULL"`
	Payment       Payment
}

type BookingReq struct {
	GamesID       uint   `json:"games_id"`
	ListTopUpId   uint   `json:"list_top_up_id"`
	UserId        string `json:"user_id"`
	ServerId      string `json:"server_id"`
	PaymentMethod string `json:"payment_method"`
	VirtualAcc    string `json:"virtual_acc"`
	NoWa          string `json:"no_wa"`
	NameAcc       string `json:"name_acc"`
}
