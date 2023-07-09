package main

import (
	"TopUpWEb/database"
	"TopUpWEb/handler"
	"TopUpWEb/initializers"
	"TopUpWEb/repository"
	"TopUpWEb/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
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

func AdminHandler(db *gorm.DB) *handler.AdminHandler {
	adminRepo := repository.NewAdminRepo(db)
	adminServ := service.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminServ)
	return adminHandler
}

func main() {
	//err := godotenv.Load()
	//
	//if err != nil {
	//	log.Fatalln("Failed to load env file")
	//}
	db := database.InitDB()
	gameHandler := GameHandler(db)
	bookingHandler := BookingHandler(db)
	paymentHandler := PaymentHandler(db)
	//handler.CreateAdmin(db)
	adminHandler := AdminHandler(db)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With, name")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "succes",
		})
	})

	r.GET("/games", gameHandler.GetAllGames)
	r.GET("/games/:id", gameHandler.GetGamebyID)
	r.GET("/booking-details", bookingHandler.ShowLatestBooking)
	r.GET("/orders-list", paymentHandler.ShowPaidOrder)
	r.GET("/orders-list/:id", paymentHandler.ShowOrderByIdAdmin)
	r.GET("/order-status", paymentHandler.ShowOrderByIdUser)
	r.GET("/done-orders", paymentHandler.ShowDoneOrder)
	r.GET("/process-orders", paymentHandler.ShowProcessOrder)
	r.GET("/latest-order", paymentHandler.ShowLatestOrder)
	r.GET("/orders-by-name", paymentHandler.FindOrderByGame)
	r.GET("/History", paymentHandler.ShowHistory)
	r.POST("/booking/:id", bookingHandler.CreateBooking)
	r.POST("/order", paymentHandler.CreateOrder)
	r.POST("/Login", adminHandler.Login)
	r.PATCH("/order/pay/:id", paymentHandler.Payment)
	r.PATCH("/order/confirm-order/:id", paymentHandler.ConfirmOrder)
	r.PATCH("/change-pass", adminHandler.ChangePass)
	r.PATCH("/change-uname", adminHandler.ChangeUname)
	r.DELETE("/admin-delete/:id", paymentHandler.DeleteOrderAdmin)
	r.DELETE("/user-delete", paymentHandler.DeleteOrderUser)

	r.Run()
}
