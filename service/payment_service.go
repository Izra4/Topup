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
