package routes

import (
	"gading-go-challenge/controllers"
	"gading-go-challenge/repositories"
	"gading-go-challenge/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteSession(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	// REPOSITORIES
	orderRepo := repositories.NewOrderRepository(db)
	itemRepo := repositories.NewItemRepository(db)
	// SERVICES
	orderService := services.NewOrderService(orderRepo, itemRepo)
	// CONTROLLERS
	orderController := controllers.NewOrderController(orderService)


	order := router.Group("orders")
	{
		order.POST("/", orderController.CreateOrder)
		order.GET("/", orderController.GetOrders)
		order.GET("/:id", orderController.GetOrder)
		order.DELETE("/:id", orderController.DeleteOrder)
		order.PUT("/:id", orderController.UpdateOrder)
	}

	return router

}
