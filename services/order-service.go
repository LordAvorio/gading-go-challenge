package services

import (
	"gading-go-challenge/models"
	"gading-go-challenge/repositories"
	"time"
)

type OrderService struct {
	orderRepo repositories.IOrderRepository
	itemRepo  repositories.IItemRepository
}

type IOrderService interface {
	CreateOrder(models.Order) (*models.OrderResponse, error)
	GetOrders() (*[]models.OrderResponse, error)
	GetOrder(int) (*models.OrderResponse, error)
	DeleteOrder(int) error
	UpdateOrder(models.Order) (*models.OrderResponse, error)
}

func NewOrderService(orderRepo repositories.IOrderRepository, itemRepo repositories.IItemRepository) *OrderService {
	orderService := OrderService{}
	orderService.orderRepo = orderRepo
	orderService.itemRepo = itemRepo
	return &orderService
}

func (s *OrderService) CreateOrder(data models.Order) (*models.OrderResponse, error) {

	dataOrder := models.Order{
		CustomerName: data.CustomerName,
		OrderedAt:    data.OrderedAt,
	}

	dataItems := []models.Item{}

	resultOrder, err := s.orderRepo.CreateOrder(&dataOrder)
	if err != nil {
		return nil, err
	}

	for _, valueItem := range data.Items {
		item := models.Item{
			Name:        valueItem.Name,
			Description: valueItem.Description,
			Quantity:    valueItem.Quantity,
			OrderId:     resultOrder.ID,
		}
		dataItems = append(dataItems, item)
	}

	resultItem, err := s.itemRepo.CreateItem(&dataItems)
	if err != nil {
		return nil, err
	}

	responseItems := []models.ItemResponse{}

	for _, itemValue := range *resultItem {
		responseItem := models.ItemResponse{
			Name:        itemValue.Name,
			Description: itemValue.Description,
			Quantity:    itemValue.Quantity,
		}
		responseItems = append(responseItems, responseItem)
	}

	responseOrder := models.OrderResponse{
		CustomerName: resultOrder.CustomerName,
		OrderedAt:    resultOrder.OrderedAt,
		Items:        responseItems,
	}

	return &responseOrder, nil
}

func (s *OrderService) GetOrders() (*[]models.OrderResponse, error) {

	resultOrders, err := s.orderRepo.GetOrders()
	if err != nil {
		return nil, err
	}

	responseOrders := []models.OrderResponse{}

	for _, valueOrder := range *resultOrders {
		responseItems := []models.ItemResponse{}

		for _, valueItem := range valueOrder.Items {
			responseItem := models.ItemResponse{
				Name:        valueItem.Name,
				Description: valueItem.Description,
				Quantity:    valueItem.Quantity,
			}
			responseItems = append(responseItems, responseItem)
		}

		responseOrder := models.OrderResponse{
			CustomerName: valueOrder.CustomerName,
			OrderedAt:    valueOrder.OrderedAt,
			Items:        responseItems,
		}

		responseOrders = append(responseOrders, responseOrder)
	}

	return &responseOrders, nil
}

func (s *OrderService) GetOrder(id int) (*models.OrderResponse, error) {

	resultOrder, err := s.orderRepo.GetOrder(id)
	if err != nil {
		return nil, err
	}

	responseItems := []models.ItemResponse{}

	for _, valueItem := range resultOrder.Items {
		responseItem := models.ItemResponse{
			Name:        valueItem.Name,
			Description: valueItem.Description,
			Quantity:    valueItem.Quantity,
		}
		responseItems = append(responseItems, responseItem)
	}

	responseOrder := models.OrderResponse{
		CustomerName: resultOrder.CustomerName,
		OrderedAt:    resultOrder.OrderedAt,
		Items:        responseItems,
	}

	return &responseOrder, nil
}

func (s *OrderService) DeleteOrder(id int) error {

	err := s.orderRepo.DeleteOrder(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) UpdateOrder(data models.Order) (*models.OrderResponse, error) {

	timeNow := time.Now()

	dataUpdateOrder := models.Order{
		ID:           data.ID,
		CustomerName: data.CustomerName,
		OrderedAt:    data.OrderedAt,
		UpdateAt:     &timeNow,
	}

	resultOrder, err := s.orderRepo.UpdateOrder(&dataUpdateOrder)
	if err != nil {
		return nil, err
	}

	responseItem := []models.ItemResponse{}

	for _, itemValue := range data.Items {
		itemValue.OrderId = resultOrder.ID
		itemValue.UpdateAt = &timeNow
		
		resultItem, err := s.itemRepo.UpdateItem(&itemValue)

		if err != nil {
			return nil, err
		}

		item := models.ItemResponse{
			Name:        resultItem.Name,
			Description: resultItem.Description,
			Quantity:    resultItem.Quantity,
		}

		responseItem = append(responseItem, item)
	}

	responseOrder := models.OrderResponse{
		CustomerName: resultOrder.CustomerName,
		OrderedAt:    resultOrder.OrderedAt,
		Items:        responseItem,
	}

	return &responseOrder, nil

}
