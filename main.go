package main

import (
	"TopUpWEb/database"
	"TopUpWEb/handler"
	"TopUpWEb/initializers"
	"TopUpWEb/repository"
	"TopUpWEb/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	fmt.Println("Hello")
	db := database.InitDB()
	gameHandler := GameHandler(db)
	bookingHandler := BookingHandler(db)
	paymentHandler := PaymentHandler(db)

	r := gin.Default()
	r.GET("/games", gameHandler.GetAllGames)
	r.GET("/games/:id", gameHandler.GetGamebyID)
	r.GET("/booking-details", bookingHandler.ShowLatestBooking)
	r.POST("/booking/:id", bookingHandler.CreateBooking)
	r.POST("/order", paymentHandler.CreateOrder)
	r.Run()
}
