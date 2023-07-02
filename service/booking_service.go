package service

import (
	"TopUpWEb/entity"
	"TopUpWEb/repository"
	"gorm.io/gorm"
)

type BookingService interface {
	Create(book entity.BookingReq) (entity.Booking, error)
	ShowLatest() (entity.Booking, error)
}

type bookingService struct {
	bookRepo repository.BookingRepository
}

func NewBookingService(bookRepo repository.BookingRepository) *bookingService {
	return &bookingService{bookRepo}
}

func (bs *bookingService) Create(book entity.BookingReq) (entity.Booking, error) {
	booking := entity.Booking{
		Model:         gorm.Model{},
		GamesID:       book.GamesID,
		ListTopUpId:   book.ListTopUpId,
		UserId:        book.UserId,
		ServerId:      book.ServerId,
		PaymentMethod: book.PaymentMethod,
		VirtualAcc:    book.VirtualAcc,
		NoWA:          book.NoWa,
		NameAcc:       book.NameAcc,
	}
	createdBooking, err := bs.bookRepo.Create(booking)
	if err != nil {
		return entity.Booking{}, err
	}

	return createdBooking, nil
}

func (bs *bookingService) ShowLatest() (entity.Booking, error) {
	return bs.bookRepo.ShowLatest()
}
