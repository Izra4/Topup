package handler

import (
	"TopUpWEb/database"
	"TopUpWEb/entity"
	"TopUpWEb/sdk"
	"TopUpWEb/service"
	"fmt"
	"github.com/gin-gonic/gin"
	storage_go "github.com/supabase-community/storage-go"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
	idGame := game.ID
	userID := latestBooking.UserId
	fmt.Println(idGame)
	if idGame == 1 {
		userID = userID + "(" + latestBooking.ServerId + ")"
	}
	results := entity.PaymentReq{
		ID:                generateOrderID(),
		Name:              game.Nama,
		JenisPaket:        topUp.JenisPaket,
		UserId:            userID,
		PaymentMethod:     latestBooking.PaymentMethod,
		NomorVA:           latestBooking.VirtualAcc,
		NameAcc:           latestBooking.NameAcc,
		PaymentStatus:     false,
		TransactionStatus: "Belum di Proses",
		PaymentLink:       "",
		BookingId:         latestBooking.ID,
	}
	createdOrder, err := ph.PaymentService.Create(results)
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
	Status         string `json:"status"`
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
			Status:         order.TransactionStatus,
		}
		orders = append(orders, paid)
	}

	sdk.Success(c, 200, "Success to get data", orders)
}

func (ph *PaymentHandler) ShowDoneOrder(c *gin.Context) {
	doneOrder, err := ph.PaymentService.GetOrderByDoneStatus()
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to load", err)
		return
	}
	if doneOrder == nil {
		sdk.Success(c, 200, "No orders done", doneOrder)
		return
	}
	sdk.Success(c, 200, "Data found", doneOrder)
}

func (ph *PaymentHandler) ShowProcessOrder(c *gin.Context) {
	processOrder, err := ph.PaymentService.GetOrderByProcessStatus()
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to load", err)
		return
	}
	sdk.Success(c, 200, "Data found", processOrder)
}

func (ph *PaymentHandler) ShowOrderByIdAdmin(c *gin.Context) {
	idOrder := c.Param("id")
	fmt.Println(idOrder)
	order, err := ph.PaymentService.GetById("#" + idOrder)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to load data", err)
		return
	}
	client := getClient()
	link := client.GetPublicUrl("Link_Bayar", order.PaymentLink)
	results := entity.PaymentRes{
		ID:                order.ID,
		Created_time:      order.CreatedAt,
		Name:              order.Name,
		Jenis_paket:       order.JenisPaket,
		UserId:            order.UserId,
		PaymentMethod:     order.PaymentMethod,
		NomorVA:           order.NomorVA,
		NameAcc:           order.NameAcc,
		PaymentStatus:     order.PaymentStatus,
		TransactionStatus: order.TransactionStatus,
		PaymentLink:       link.SignedURL,
	}
	sdk.Success(c, 200, "Data loaded", results)
}

func (ph *PaymentHandler) ShowOrderByIdUser(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		sdk.Fail(c, 400, "Input the id")
		return
	}

	results, err := ph.PaymentService.GetById(id)
	if err != nil {
		sdk.FailOrError(c, 400, "Order with id: "+id+"isn't exist", err)
		return
	}
	linkStr := ""
	if results.PaymentLink != "" {
		client := getClient()
		link := client.GetPublicUrl("Link_Bayar", results.PaymentLink)
		linkStr = link.SignedURL
	}

	res := entity.PaymentRes{
		ID:                results.ID,
		Created_time:      results.CreatedAt,
		Name:              results.Name,
		Jenis_paket:       results.JenisPaket,
		UserId:            results.UserId,
		PaymentMethod:     results.PaymentMethod,
		NomorVA:           results.NomorVA,
		NameAcc:           results.NameAcc,
		PaymentStatus:     results.PaymentStatus,
		TransactionStatus: results.TransactionStatus,
		PaymentLink:       linkStr,
	}
	sdk.Success(c, 200, "Data Found", res)
}

func (ph *PaymentHandler) Payment(c *gin.Context) {
	id := c.Param("id")
	id = "#" + id

	isPaid, ok := strconv.ParseBool(c.PostForm("is_paid"))
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert boolean", ok)
		return
	}
	_, err := ph.PaymentService.GetById(id)
	if err != nil {
		sdk.FailOrError(c, 500, "Data not found", err)
		return
	}
	transacStatus := c.PostForm("status")
	if transacStatus == "" {
		sdk.Fail(c, 400, "Failed to get new transaction status")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		sdk.FailOrError(c, http.StatusBadRequest, "Failed to get file", err)
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to open file", err)
		return
	}
	defer fileContent.Close()

	client := getClient()
	fileName := randString()
	resp := client.UploadFile("Link_Bayar", fileName, fileContent)
	fmt.Println(resp)
	res := result{
		Id:          id,
		IsPaid:      isPaid,
		TransStatus: transacStatus,
		Link:        fileName,
	}
	err = ph.PaymentService.UpdatePayment(id, isPaid, transacStatus, fileName)
	if err != nil {
		sdk.FailOrError(c, 500, "payment failed", err)
		return
	}
	sdk.Success(c, 200, "Succes to pay", res)
}

type result struct {
	Id          string `json:"id"`
	IsPaid      bool   `json:"is_paid"`
	TransStatus string `json:"trans_status"`
	Link        string `json:"link"`
}

func (ph *PaymentHandler) ConfirmOrder(c *gin.Context) {
	id := c.Param("id")
	id = "#" + id

	status := c.PostForm("status")
	if status == "" {
		sdk.Fail(c, 500, "Failed to get status")
		return
	}
	if err := ph.PaymentService.OrderConfirm(id, status); err != nil {
		sdk.FailOrError(c, 500, "Failed to update status", err)
		return
	}
	res := orderResult{
		ID:     id,
		Status: status,
	}
	sdk.Success(c, 200, "Order success", res)
}

func (ph *PaymentHandler) DeleteOrderAdmin(c *gin.Context) {
	id := c.Param("id")
	id = "#" + id
	data, err := ph.PaymentService.DeleteOrder(id)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to delete", err)
		return
	}
	sdk.Success(c, 200, "Order with id "+id+" deleted", data)
}
func (ph *PaymentHandler) DeleteOrderUser(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		sdk.Fail(c, 400, "Input the id")
		return
	}
	data, err := ph.PaymentService.DeleteOrder(id)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to delete", err)
		return
	}
	sdk.Success(c, 200, "Order with id: "+id+" deleted", data)
}

type orderResult struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (ph *PaymentHandler) ShowLatestOrder(c *gin.Context) {
	data, err := ph.PaymentService.ShowLatestOrder()
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}
	booking, err := ph.BookingService.FindById(data.BookingId)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}
	var topUp entity.ListTopUp
	if err = database.DB.Where("id = ?", booking.ListTopUpId).Find(&topUp).Error; err != nil {
		sdk.FailOrError(c, 500, "Failed to get top up data", err)
		return
	}
	res := entity.ShowOrder{
		ID:                data.ID,
		Name:              data.Name,
		Jenis_paket:       data.JenisPaket,
		UserId:            data.UserId,
		PaymentMethod:     data.PaymentMethod,
		NomorVA:           data.NomorVA,
		NameAcc:           data.NameAcc,
		PaymentStatus:     false,
		TransactionStatus: data.TransactionStatus,
		Harga:             topUp.Harga,
	}
	sdk.Success(c, 200, "Success", res)
}

func (ph *PaymentHandler) FindOrderByGame(c *gin.Context) {
	gameName := c.PostForm("name")
	if gameName == "" {
		sdk.Fail(c, 500, "failed to get game name")
		return
	}

	if gameName == "1" {
		gameName = "Mobile Legends"
	} else if gameName == "2" {
		gameName = "PUBG Mobile"
	} else if gameName == "3" {
		gameName = "Valorant"
	}

	payments, err := ph.PaymentService.FindOrderByGame(gameName)
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to find orders by game", err)
		return
	}

	sdk.Success(c, http.StatusOK, "Orders found by game", payments)
}

func generateOrderID() string {
	// Generate random number with 8 characters
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)
	randNum := fmt.Sprintf("%08d", randGen.Intn(100000000))

	orderID := "#INV" + randNum
	return orderID
}

func getClient() *storage_go.Client {
	client := storage_go.NewClient(os.Getenv("PROJECT_URL"), os.Getenv("PROJECT_API"), nil)
	return client
}

func randString() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 8)
	for i := range b {
		b[i] = chars[randSeed.Intn(len(chars))]
	}

	return string(b)
}
