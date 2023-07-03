package repository

import (
	"TopUpWEb/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(order entity.Payment) (entity.Payment, error)
	ShowPaidOrder() ([]entity.Payment, error)
	GetById(id string) (entity.Payment, error)
	UpdatePayment(id string, paymentStatus bool, transacStatus string, paymentLink string) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db}
}

func (pr *paymentRepository) Create(order entity.Payment) (entity.Payment, error) {
	if err := pr.db.Create(&order).Error; err != nil {
		return entity.Payment{}, err
	}
	return order, nil
}

func (pr *paymentRepository) ShowPaidOrder() ([]entity.Payment, error) {
	var result []entity.Payment
	if err := pr.db.Where("is_paid = ?", true).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (pr *paymentRepository) GetById(id string) (entity.Payment, error) {
	var result entity.Payment
	if err := pr.db.First(&result, "id = ?", id).Error; err != nil {
		return entity.Payment{}, err
	}
	return result, nil
}

func (pr *paymentRepository) UpdatePayment(id string, paymentStatus bool, transacStatus string, link string) error {
	var order entity.Payment
	if err := pr.db.Model(&order).Where("id = ?", id).Updates(map[string]interface{}{
		"is_paid":            paymentStatus,
		"transaction_status": transacStatus,
		"payment_link":       link,
	}).Error; err != nil {
		return err
	}
	return nil
}
