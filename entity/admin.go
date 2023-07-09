package entity

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"NOT NULL"`
	Password string `gorm:"NOT NULL"`
}

type AdminReq struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}
