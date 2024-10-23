package server

import (
	"htmx-go/internals/db"
	"htmx-go/internals/handlers"
	"htmx-go/internals/repository"
	"htmx-go/internals/services"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	router := gin.Default()

	dbConn := db.Connect()

	orderRepo := repository.NewOrderRepository(dbConn)
	userRepo := repository.NewUserRepository(dbConn)

	orderService := services.NewOrderService(orderRepo)
	userService := services.NewUserService(userRepo)

	indexHandler := handlers.NewIndexHandler(orderService)
	orderHandler := handlers.NewOrderHandler(orderService)
	userHandler := handlers.NewUserHandler(userService)

	router.GET("/", indexHandler.IndexPageHandler)
	router.POST("/orders", orderHandler.CreateOrder)
	router.GET("/orders/:id", orderHandler.GetOrder)
	router.PUT("/orders/:id", orderHandler.UpdateOrder)
	router.DELETE("/orders/:id", orderHandler.DeleteOrder)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	return router
}
