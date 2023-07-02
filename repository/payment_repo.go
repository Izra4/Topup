package repository

import (
	"TopUpWEb/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(order entity.Payment) (entity.Payment, error)
	ShowPaidOrder() ([]entity.Payment, error)
	GetById(id uint) (entity.Payment, error)
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

func (pr *paymentRepository) GetById(id uint) (entity.Payment, error) {
	var result entity.Payment
	if err := pr.db.First(&result, id).Error; err != nil {
		return entity.Payment{}, err
	}
	return result, nil
}
