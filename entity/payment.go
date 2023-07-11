package entity

import (
	"time"
)

type Payment struct {
	ID                string    `gorm:"primarykey;type:VARCHAR(12)" json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Name              string    `gorm:"type:VARCHAR(30);NOT NULL" json:"name"`
	JenisPaket        string    `gorm:"type:VARCHAR(20);NOT NULL" json:"jenis_paket"`
	UserId            string    `gorm:"type:VARCHAR(30);NOT NULL" json:"user_id"`
	PaymentMethod     string    `gorm:"type:VARCHAR(30);NOT NULL" json:"payment_method"`
	NomorVA           string    `gorm:"type:VARCHAR(20);NOT NULL" json:"nomor_va"`
	NameAcc           string    `gorm:"type:VARCHAR(30);NOT NULL" json:"name_acc"`
	PaymentStatus     bool      `gorm:"column:is_paid;default:false" json:"payment_status"`
	TransactionStatus string    `gorm:"type:VARCHAR(50);NOT NULL" json:"transaction_status"`
	PaymentLink       string    `gorm:"type:TEXT;NOT NULL" json:"payment_link"`
	BookingId         uint      `gorm:"unique" json:"booking_id"`
}

type PaymentRes struct {
	ID                string    `json:"id"`
	Created_time      time.Time `json:"created_time"`
	Name              string    `json:"name"`
	Jenis_paket       string    `json:"jenis_paket"`
	Harga             string    `json:"harga"`
	UserId            string    `json:"user_id"`
	PaymentMethod     string    `json:"payment_method"`
	NomorVA           string    `json:"nomor_va"`
	NameAcc           string    `json:"name_acc"`
	PaymentStatus     bool      `json:"payment_status"`
	TransactionStatus string    `json:"transaction_status"`
	PaymentLink       string    `json:"payment_link"`
}

type PaymentReq struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	JenisPaket        string `json:"jenis_paket"`
	UserId            string `json:"user_id"`
	PaymentMethod     string `json:"payment_method"`
	NomorVA           string `json:"nomor_va"`
	NameAcc           string `json:"name_acc"`
	PaymentStatus     bool   `json:"payment_status"`
	TransactionStatus string `json:"transaction_status"`
	PaymentLink       string `json:"payment_link"`
	BookingId         uint   `json:"booking_id"`
}

type ShowOrder struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Jenis_paket       string `json:"jenis_paket"`
	UserId            string `json:"user_id"`
	PaymentMethod     string `json:"payment_method"`
	NomorVA           string `json:"nomor_va"`
	NameAcc           string `json:"name_acc"`
	PaymentStatus     bool   `json:"payment_status"`
	TransactionStatus string `json:"transaction_status"`
	Harga             string `json:"harga"`
}
