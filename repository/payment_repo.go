package repository

import (
	"TopUpWEb/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(order entity.Payment) (entity.Payment, error)
	ShowPaidOrder() ([]entity.Payment, error)
	GetById(id string) (entity.Payment, error)
	GetOrderByDoneStatus() ([]entity.Payment, error)
	GetOrderByProcessStatus() ([]entity.Payment, error)
	UpdatePayment(id string, paymentStatus bool, transacStatus string, paymentLink string) error
	OrderConfirm(id string, transacStatus string) error
	DeleteOrder(id string) (entity.Payment, error)
	ShowLatestOrder() (entity.Payment, error)
	FindOrderByGame(name string) ([]entity.Payment, error)
	OrderHistory() ([]entity.Payment, error)
	GetOrderByFailStatus() ([]entity.Payment, error)
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

func (pr *paymentRepository) OrderConfirm(id string, transacStatus string) error {
	var order entity.Payment
	if err := pr.db.Model(&order).Where("id = ?", id).Updates(map[string]interface{}{
		"transaction_status": transacStatus,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (pr *paymentRepository) DeleteOrder(id string) (entity.Payment, error) {
	var order entity.Payment
	if err := pr.db.First(&order, "id = ?", id).Error; err != nil {
		return order, err
	}
	if err := pr.db.Delete(&order).Error; err != nil {
		return order, err
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

func (pr *paymentRepository) GetOrderByDoneStatus() ([]entity.Payment, error) {
	var result []entity.Payment
	if err := pr.db.Where("transaction_status = ?", "Selesai").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (pr *paymentRepository) GetOrderByProcessStatus() ([]entity.Payment, error) {
	var result []entity.Payment
	if err := pr.db.Where("transaction_status = ?", "Sedang Proses").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (pr *paymentRepository) GetOrderByFailStatus() ([]entity.Payment, error) {
	var result []entity.Payment
	if err := pr.db.Where("transaction_status = ?", "Pembayaran invalid").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (pr *paymentRepository) ShowLatestOrder() (entity.Payment, error) {
	var data entity.Payment
	if err := pr.db.Order("created_at desc").First(&data).Error; err != nil {
		return entity.Payment{}, err
	}
	return data, nil
}

func (pr *paymentRepository) FindOrderByGame(name string) ([]entity.Payment, error) {
	var payments []entity.Payment
	if err := pr.db.Where("name = ? AND is_paid = ?", name, true).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (pr *paymentRepository) OrderHistory() ([]entity.Payment, error) {
	var datas []entity.Payment
	if err := pr.db.Where("transaction_status = ?", "Selesai").
		Order("updated_at desc").Limit(10).Find(&datas).Error; err != nil {
		return nil, err
	}
	return datas, nil
}
