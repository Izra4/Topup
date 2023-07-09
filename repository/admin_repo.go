package repository

import (
	"TopUpWEb/entity"
	"gorm.io/gorm"
)

type AdminRepo interface {
	Create(admin entity.Admin) (entity.Admin, error)
	UpdatePass(id uint, pass string) error
	UpdateName(id uint, name string) error
}

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *adminRepo {
	return &adminRepo{db}
}

func (ar *adminRepo) Create(admin entity.Admin) (entity.Admin, error) {
	if err := ar.db.Create(&admin).Error; err != nil {
		return entity.Admin{}, err
	}
	return admin, nil
}

func (ar *adminRepo) UpdatePass(id uint, pass string) error {
	if err := ar.db.Model(&entity.Admin{}).Where("id = ?", id).Update("password", pass).Error; err != nil {
		return err
	}
	return nil
}

func (ar *adminRepo) UpdateName(id uint, name string) error {
	if err := ar.db.Model(&entity.Admin{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}
