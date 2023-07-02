package handler

import (
	"TopUpWEb/database"
	"TopUpWEb/entity"
	"TopUpWEb/sdk"
	"TopUpWEb/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookingHandler struct {
	bookingService service.BookingService
}

func NewBookingHandler(bookingService service.BookingService) *BookingHandler {
	return &BookingHandler{bookingService}
}

func (bh *BookingHandler) CreateBooking(c *gin.Context) {
	strGamesId := c.Param("id")
	serverId := ""
	//GET DATA
	gamesId, ok := strconv.Atoi(strGamesId)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to Convert", ok)
		return
	}
	strIdTopUp := c.PostForm("IdTopUp")
	if strIdTopUp == "" {
		sdk.Fail(c, 400, "Please select the type of topup you want to buy")
		return
	}
	idTopUp, ok := strconv.Atoi(strIdTopUp)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to Convert", ok)
		return
	}

	paymentMethod := c.PostForm("payment")
	if paymentMethod == "" {
		sdk.Fail(c, 400, "Please select your payment method")
		return
	}
	VA := c.PostForm("VirtualAcc")
	if VA == "" {
		sdk.Fail(c, 400, "Failed to load Virtual Account")
		return
	}
	NoWA := c.PostForm("nomor")
	if NoWA == "" {
		sdk.Fail(c, 400, "Please insert your number")
		return
	}

	userId := c.PostForm("idUser")
	if userId == "" {
		sdk.Fail(c, 400, "Please insert user Id")
		return
	}
	nameAcc := c.PostForm("name_acc")
	if nameAcc == "" {
		sdk.Fail(c, 500, "Acc Name is Empty")
		return
	}
	if gamesId == 1 {
		serverId = c.PostForm("ServerId")
		if serverId == "" {
			sdk.Fail(c, 400, "Please insert server Id")
			return
		}
	}

	//CREATE DATA
	booking := entity.BookingReq{
		GamesID:       uint(gamesId),
		ListTopUpId:   uint(idTopUp),
		UserId:        userId,
		ServerId:      serverId,
		PaymentMethod: paymentMethod,
		VirtualAcc:    VA,
		NoWa:          NoWA,
		NameAcc:       nameAcc,
	}

	createBooking, err := bh.bookingService.Create(booking)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to Create Booking Data", err)
		return
	}
	fmt.Println(nameAcc)
	sdk.Success(c, 200, "Booking data created", createBooking)
}

type mlResult struct {
	GameName         string `json:"game_name"`
	UserId           string `json:"user_id"`
	ServerId         string `json:"server_id"`
	Data             string `json:"data"`
	Harga            string `json:"harga"`
	MetodePembayaran string
}

type Result struct {
	GameName         string `json:"game_name"`
	UserId           string `json:"user_id"`
	Data             string `json:"data"`
	Harga            string `json:"harga"`
	MetodePembayaran string
}

func (bh *BookingHandler) ShowLatestBooking(c *gin.Context) {
	data, err := bh.bookingService.ShowLatest()
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}
	var game entity.Games
	if err = database.DB.Where("id = ?", data.GamesID).Find(&game).Error; err != nil {
		sdk.FailOrError(c, 500, "Failed to get game data", err)
		return
	}

	var topUp entity.ListTopUp
	if err = database.DB.Where("id = ?", data.ListTopUpId).Find(&topUp).Error; err != nil {
		sdk.FailOrError(c, 500, "Failed to get top up data", err)
		return
	}
	if data.GamesID == 1 {
		result := mlResult{
			GameName:         game.Nama,
			UserId:           data.UserId,
			ServerId:         data.ServerId,
			Data:             topUp.JenisPaket,
			Harga:            topUp.Harga,
			MetodePembayaran: data.PaymentMethod,
		}
		sdk.Success(c, 200, "Data Found", result)
		return
	}
	result := Result{
		GameName:         game.Nama,
		UserId:           data.UserId,
		Data:             topUp.JenisPaket,
		Harga:            topUp.Harga,
		MetodePembayaran: data.PaymentMethod,
	}
	sdk.Success(c, 200, "Data Found", result)

}
