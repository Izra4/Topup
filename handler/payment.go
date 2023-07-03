package handler

import (
	"TopUpWEb/database"
	"TopUpWEb/entity"
	"TopUpWEb/sdk"
	"TopUpWEb/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type PaymentHandler struct {
	PaymentService service.PaymentService
	BookingService service.BookingService
	GameService    service.GameService
}

func NewPaymentHandler(paymentService service.PaymentService, bookingService service.BookingService, gameService service.GameService) *PaymentHandler {
	return &PaymentHandler{
		PaymentService: paymentService,
		BookingService: bookingService,
		GameService:    gameService,
	}
}

func (ph *PaymentHandler) CreateOrder(c *gin.Context) {
	latestBooking, err := ph.BookingService.ShowLatest()
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get latest booking data", err)
		return
	}

	game, err := ph.GameService.FindbyId(latestBooking.GamesID)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get game data", err)
		return
	}
	var topUp entity.ListTopUp
	if err = database.DB.First(&topUp, latestBooking.ListTopUpId).Error; err != nil {
		sdk.FailOrError(c, 500, "Failed to get topup data", err)
		return
	}

	result := entity.PaymentReq{
		ID:                generateOrderID(),
		Name:              game.Nama,
		JenisPaket:        topUp.JenisPaket,
		UserId:            latestBooking.UserId,
		PaymentMethod:     latestBooking.PaymentMethod,
		NomorVA:           latestBooking.VirtualAcc,
		NameAcc:           latestBooking.NameAcc,
		PaymentStatus:     false,
		TransactionStatus: "Belum di Proses",
		PaymentLink:       "",
		BookingId:         latestBooking.ID,
	}
	createdOrder, err := ph.PaymentService.Create(result)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to create order", err)
		return
	}
	sdk.Success(c, 200, "Success to create order data", createdOrder)
}

type orderList struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	JenisPaket     string `json:"jenis_paket"`
	PaymentMeethod string `json:"payment_meethod"`
}

func (ph *PaymentHandler) ShowPaidOrder(c *gin.Context) {
	paidOrder, err := ph.PaymentService.ShowPaidOrder()
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to load", err)
		return
	}
	if paidOrder == nil {
		sdk.Success(c, 200, "No orders paid", paidOrder)
		return
	}

	var orders []orderList
	for _, order := range paidOrder {
		paid := orderList{
			Id:             order.ID,
			Name:           order.Name,
			JenisPaket:     order.JenisPaket,
			PaymentMeethod: order.PaymentMethod,
		}
		orders = append(orders, paid)
	}

	sdk.Success(c, 200, "Success to get data", orders)
}

func (ph *PaymentHandler) ShowOrderById(c *gin.Context) {
	idOrder := c.Param("id")
	fmt.Println(idOrder)
	order, err := ph.PaymentService.GetById("#" + idOrder)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to load data", err)
		return
	}
	sdk.Success(c, 200, "Data loaded", order)
}

func generateOrderID() string {
	// Generate random number with 8 characters
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)
	randNum := fmt.Sprintf("%08d", randGen.Intn(100000000))

	orderID := "#INV" + randNum
	return orderID
}
