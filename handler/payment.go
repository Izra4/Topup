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

type idReq struct {
	Id string `json:"id" binding:"required"`
}

func (ph *PaymentHandler) ShowOrderByIdUser(c *gin.Context) {
	var req idReq
	if err := c.ShouldBindJSON(&req); err != nil {
		sdk.FailOrError(c, 400, "Please input the id that you want to check", err)
	}

	results, err := ph.PaymentService.GetById(req.Id)
	if err != nil {
		sdk.FailOrError(c, 400, "Order with id: "+req.Id+"isn't exist", err)
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
	var req idReq
	if err := c.ShouldBindJSON(&req); err != nil {
		sdk.FailOrError(c, 400, "Please input the id that you want to delete", err)
		return
	}
	data, err := ph.PaymentService.DeleteOrder(req.Id)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to delete", err)
		return
	}
	sdk.Success(c, 200, "Order with id: "+req.Id+" deleted", data)
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
	sdk.Success(c, 200, "Success", data)
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
