package repositories

import (
	"gading-go-challenge/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

type IOrderRepository interface {
	CreateOrder(*models.Order) (*models.Order, error)
	GetOrders() (*[]models.Order, error)
	GetOrder(int) (*models.Order, error)
	UpdateOrder(*models.Order) (*models.Order, error)
	DeleteOrder(int) error
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	orderRepository := OrderRepository{}
	orderRepository.db = db
	return &orderRepository
}

func (r *OrderRepository) CreateOrder(order *models.Order) (*models.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetOrders() (*[]models.Order, error) {
	orders := []models.Order{}

	err := r.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (r *OrderRepository) GetOrder(id int) (*models.Order, error) {
	order := models.Order{}

	err := r.db.Preload("Items").Where("id = ?", id).Take(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) DeleteOrder(id int) error {
	order := models.Order{}

	err := r.db.Preload("Items").Where("id = ?", id).Delete(&order).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) UpdateOrder(order *models.Order) (*models.Order, error) {
	err := r.db.Model(&order).Where("id = ?", order.ID).Updates(order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}


