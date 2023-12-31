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
	tokenStr, err := sdk.Token(admin)
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "create token failed", err)
		return
	}
	sdk.Success(c, 200, "Login Success", map[string]string{"token": tokenStr})
}

func (ah *AdminHandler) ChangePass(c *gin.Context) {
	id, _ := c.Get("admin")
	claims := id.(entity.AdminClaims)

	oldPass := c.PostForm("oldPass")
	if oldPass == "" {
		sdk.Fail(c, 400, "Old pass value is empty")
		return
	}
	newPass := c.PostForm("newPass")
	if newPass == "" {
		sdk.Fail(c, 400, "New pass value is empty")
		return
	}
	var data entity.Admin
	if err := database.DB.First(&data, claims.ID).Error; err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}
	if err := ValidateHash(data.Password, oldPass); err != nil {
		sdk.FailOrError(c, 400, "Invalid pass", err)
		return
	}
	hashedPass, err := Hashing(newPass)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to hash", err)
		return
	}
	if err = ah.adminService.UpdateAdminPassword(claims.ID, hashedPass); err != nil {
		sdk.FailOrError(c, 500, "Failed to update", err)
		return
	}
	sdk.Success(c, 200, "Password changed", nil)
}

func (ah *AdminHandler) ChangeUname(c *gin.Context) {
	id, _ := c.Get("admin")
	claims := id.(entity.AdminClaims)

	newUname := c.PostForm("uname")
	if newUname == "" {
		sdk.Fail(c, 400, "Nothing changed")
		return
	}
	var data entity.Admin
	if err := database.DB.First(&data, claims.ID).Error; err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}
	if err := ah.adminService.UpdateAdminName(claims.ID, newUname); err != nil {
		sdk.FailOrError(c, 500, "Failed to update new name", err)
		return
	}
	sdk.Success(c, 200, "Success to change name", gin.H{
		"new_uname": newUname,
	})

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
