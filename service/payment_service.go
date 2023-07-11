package service

import (
	"TopUpWEb/entity"
	"TopUpWEb/repository"
	"time"
)

type PaymentService interface {
	Create(order entity.PaymentReq) (entity.Payment, error)
	ShowPaidOrder() ([]entity.Payment, error)
	GetById(id string) (entity.Payment, error)
	UpdatePayment(paymentID string, isPaid bool, transactionStatus string, link string) error
	OrderConfirm(id string, transacStatus string) error
	DeleteOrder(id string) (entity.Payment, error)
	GetOrderByDoneStatus() ([]entity.Payment, error)
	GetOrderByProcessStatus() ([]entity.Payment, error)
	ShowLatestOrder() (entity.Payment, error)
	FindOrderByGame(name string) ([]entity.Payment, error)
	OrderHistory() ([]entity.Payment, error)
	GetOrderByFailStatus() ([]entity.Payment, error)
}

type paymentService struct {
	PaymentRepo repository.PaymentRepository
}

func NewPaymentService(PaymentRepo repository.PaymentRepository) *paymentService {
	return &paymentService{PaymentRepo}
}

func (ps *paymentService) Create(order entity.PaymentReq) (entity.Payment, error) {
	orderReq := entity.Payment{
		ID:                order.ID,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		Name:              order.Name,
		JenisPaket:        order.JenisPaket,
		UserId:            order.UserId,
		PaymentMethod:     order.PaymentMethod,
		NomorVA:           order.NomorVA,
		NameAcc:           order.NameAcc,
		PaymentStatus:     false,
		TransactionStatus: order.TransactionStatus,
		PaymentLink:       order.PaymentLink,
		BookingId:         order.BookingId,
	}

	createdOrder, err := ps.PaymentRepo.Create(orderReq)
	if err != nil {
		return entity.Payment{}, err
	}
	return createdOrder, nil
}

func (ps *paymentService) ShowPaidOrder() ([]entity.Payment, error) {
	return ps.PaymentRepo.ShowPaidOrder()
}

func (ps *paymentService) GetById(id string) (entity.Payment, error) {
	return ps.PaymentRepo.GetById(id)
}

func (ps *paymentService) UpdatePayment(paymentID string, isPaid bool, transactionStatus string, link string) error {
	return ps.PaymentRepo.UpdatePayment(paymentID, isPaid, transactionStatus, link)
}

func (ps *paymentService) OrderConfirm(id string, transacStatus string) error {
	return ps.PaymentRepo.OrderConfirm(id, transacStatus)
}

func (ps *paymentService) DeleteOrder(id string) (entity.Payment, error) {
	return ps.PaymentRepo.DeleteOrder(id)
}

func (ps *paymentService) GetOrderByProcessStatus() ([]entity.Payment, error) {
	return ps.PaymentRepo.GetOrderByProcessStatus()
}

func (ps *paymentService) GetOrderByDoneStatus() ([]entity.Payment, error) {
	return ps.PaymentRepo.GetOrderByDoneStatus()
}

func (ps *paymentService) ShowLatestOrder() (entity.Payment, error) {
	return ps.PaymentRepo.ShowLatestOrder()
}

func (ps *paymentService) FindOrderByGame(name string) ([]entity.Payment, error) {
	return ps.PaymentRepo.FindOrderByGame(name)
}

func (ps *paymentService) OrderHistory() ([]entity.Payment, error) {
	return ps.PaymentRepo.OrderHistory()
}

func (ps *paymentService) GetOrderByFailStatus() ([]entity.Payment, error) {
	return ps.PaymentRepo.GetOrderByFailStatus()
}
