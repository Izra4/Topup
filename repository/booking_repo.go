package repository

import (
	"TopUpWEb/entity"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(book entity.Booking) (entity.Booking, error)
	ShowLatest() (entity.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *bookingRepository {
	return &bookingRepository{db}
}

func (br *bookingRepository) Create(book entity.Booking) (entity.Booking, error) {
	if err := br.db.Create(&book).Error; err != nil {
		return entity.Booking{}, err
	}
	return book, nil
}

func (br *bookingRepository) ShowLatest() (entity.Booking, error) {
	var booking entity.Booking
	if err := br.db.Last(&booking).Error; err != nil {
		return entity.Booking{}, err
	}
	return booking, nil
}
