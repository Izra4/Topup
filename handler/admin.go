package handler

import (
	"TopUpWEb/database"
	"TopUpWEb/entity"
	"TopUpWEb/repository"
	"TopUpWEb/sdk"
	"TopUpWEb/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type AdminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{adminService}
}

func CreateAdmin(db *gorm.DB) {
	adminRepo := repository.NewAdminRepo(db)
	adminService := service.NewAdminService(adminRepo)
	hashedPassword, err := Hashing("admin")

	admin := entity.AdminReq{
		Name: "admin",
		Pass: hashedPassword,
	}

	createdAdmin, err := adminService.CreateAdmin(admin)
	if err != nil {
		fmt.Println("Failed to create admin:", err)
		return
	}

	fmt.Println("Admin created:", createdAdmin)
}

func (ah *AdminHandler) Login(c *gin.Context) {
	uname := c.PostForm("uname")
	if uname == "" {
		sdk.Fail(c, 400, "Your name is empty")
		return
	}
	pass := c.PostForm("pass")
	if pass == "" {
		sdk.Fail(c, 400, "Your pass is empty")
		return
	}
	var admin entity.Admin
	if err := database.DB.First(&admin, "name = ?", uname).Error; err != nil {
		sdk.Fail(c, http.StatusNotFound, "Invalid name")
		return
	}
	if err := ValidateHash(admin.Password, pass); err != nil {
		sdk.Fail(c, http.StatusNotFound, "Invalid pass")
		return
	}
	sdk.Success(c, 200, "Login Success", nil)
}

func Hashing(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), err
}

func ValidateHash(hashed, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	return err
}
