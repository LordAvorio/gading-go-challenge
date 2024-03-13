package controllers

import (
	"gading-go-challenge/models"
	"gading-go-challenge/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.IOrderService
}

func NewOrderController(orderService services.IOrderService) *OrderController {
	orderController := OrderController{}
	orderController.orderService = orderService
	return &orderController
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {

	orderRequest := models.OrderRequest{}
	contentType := ctx.Request.Header.Get("Content-Type")

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&orderRequest)
	} else {
		models.ResponseError(ctx, "Request should be form in JSON Body", http.StatusBadRequest)
		return
	}

	orderData := models.Order{
		CustomerName: orderRequest.CustomerName,
		Items:        orderRequest.Items,
		OrderedAt:    orderRequest.OrderedAt,
	}

	result, err := c.orderService.CreateOrder(orderData)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	models.ResponseSuccessCreated(ctx, result)

}

func (c *OrderController) GetOrders(ctx *gin.Context) {

	result, err := c.orderService.GetOrders()
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	models.ResponseSuccessWithData(ctx, result)
}

func (c *OrderController) GetOrder(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.orderService.GetOrder(id)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	models.ResponseSuccessWithData(ctx, result)
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.orderService.DeleteOrder(id)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	models.ResponseSuccess(ctx)
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	orderRequest := models.OrderRequest{}
	contentType := ctx.Request.Header.Get("Content-Type")

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&orderRequest)
	} else {
		models.ResponseError(ctx, "Request should be form in JSON Body", http.StatusBadRequest)
		return
	}

	orderData := models.Order{
		ID:           uint(id),
		CustomerName: orderRequest.CustomerName,
		Items:        orderRequest.Items,
		OrderedAt:    orderRequest.OrderedAt,
	}

	result, err := c.orderService.UpdateOrder(orderData)
	if err != nil {
		models.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	models.ResponseSuccessCreated(ctx, result)

}
