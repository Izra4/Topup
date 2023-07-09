package service

import (
	"TopUpWEb/entity"
	"TopUpWEb/repository"
	"gorm.io/gorm"
)

type AdminService interface {
	CreateAdmin(admin entity.AdminReq) (entity.Admin, error)
	UpdateAdminName(adminID uint, name string) error
	UpdateAdminPassword(adminID uint, password string) error
}

type adminService struct {
	adminRepo repository.AdminRepo
}

func NewAdminService(adminRepo repository.AdminRepo) *adminService {
	return &adminService{adminRepo}
}

func (as *adminService) CreateAdmin(admin entity.AdminReq) (entity.Admin, error) {
	data := entity.Admin{
		Model:    gorm.Model{},
		Name:     admin.Name,
		Password: admin.Pass,
	}
	return as.adminRepo.Create(data)
}

func (as *adminService) UpdateAdminName(adminID uint, name string) error {
	return as.adminRepo.UpdateName(adminID, name)
}

func (as *adminService) UpdateAdminPassword(adminID uint, password string) error {
	return as.adminRepo.UpdatePass(adminID, password)
}
