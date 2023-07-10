package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	gorm.Model
	Name     string `gorm:"NOT NULL"`
	Password string `gorm:"NOT NULL"`
}

type AdminReq struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}
type AdminClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func NewAdminClaims(id uint, exp time.Duration) AdminClaims {
	return AdminClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}
