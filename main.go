package main

import (
	"TopUpWEb/database"
	"TopUpWEb/handler"
	"TopUpWEb/initializers"
	"TopUpWEb/repository"
	"TopUpWEb/service"
	"fmt"
	"github.com/gin-gonic/gin"
	storage_go "github.com/supabase-community/storage-go"
	"gorm.io/gorm"
	"os"
)

func init() {
	initializers.LoadEnv()
	initializers.SyncDb()
	//handler.GamesData()
	//handler.MLTopUp()
	//handler.PUBGTopUp()
	//handler.ValorantTopUp()
}

func GameHandler(db *gorm.DB) *handler.GamesHandler {
	gameRepository := repository.NewGamesRepository(db)
	gameService := service.NewGameService(gameRepository)
	gameHandler := handler.NewGamesHandler(gameService)
	return gameHandler
}

func BookingHandler(db *gorm.DB) *handler.BookingHandler {
	bookingRepository := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepository)
	bookingHandler := handler.NewBookingHandler(bookingService)
	return bookingHandler
}

func PaymentHandler(db *gorm.DB) *handler.PaymentHandler {
	gameRepository := repository.NewGamesRepository(db)
	gameService := service.NewGameService(gameRepository)

	bookingRepository := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepository)

	paymentRepository := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService, bookingService, gameService)
	return paymentHandler
}

func main() {
	db := database.InitDB()
	gameHandler := GameHandler(db)
	bookingHandler := BookingHandler(db)
	paymentHandler := PaymentHandler(db)

	client := storage_go.NewClient(os.Getenv("PROJECT_URL"), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImdtbGNzbWl2aG1kanhpc3FjeXJ3Iiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTY4ODM3NjQ2NywiZXhwIjoyMDAzOTUyNDY3fQ.ypOZNGNqPHWd57crJPDHsClq3s5-GZ1_JnGrDJl-IZ4", nil)
	fmt.Println(client.ListBuckets())
	file, err := os.Open("C:\\Users\\INTEL\\Documents\\lol.txt")
	if err != nil {
		panic(err)
	}

	resp := client.UploadFile("Link_Bayar", "file.txt", file)
	fmt.Println("respon: ", resp)

	r := gin.Default()
	r.GET("/games", gameHandler.GetAllGames)
	r.GET("/games/:id", gameHandler.GetGamebyID)
	r.GET("/booking-details", bookingHandler.ShowLatestBooking)
	r.GET("/orders-list", paymentHandler.ShowPaidOrder)
	r.GET("/orders-list/:id", paymentHandler.ShowOrderByIdAdmin)
	r.GET("/order-status", paymentHandler.ShowOrderByIdUser)
	r.POST("/booking/:id", bookingHandler.CreateBooking)
	r.POST("/order", paymentHandler.CreateOrder)
	r.POST("/order/pay/:id", paymentHandler.Payment)
	r.Run()
}
