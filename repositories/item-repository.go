package repositories

import (
	"errors"
	"gading-go-challenge/models"

	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

type IItemRepository interface {
	CreateItem(*[]models.Item) (*[]models.Item, error)
	UpdateItem(*models.Item) (*models.Item, error)
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	itemRepository := ItemRepository{}
	itemRepository.db = db
	return &itemRepository
}

func (r *ItemRepository) CreateItem(items *[]models.Item) (*[]models.Item, error) {
	err := r.db.Create(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) UpdateItem(item *models.Item) (*models.Item, error) {

	insertNewItem := false
	dataItemDB := models.Item{}

	errCheck := r.db.Where("order_id = ? AND name = ?", item.OrderId, item.Name).First(&dataItemDB).Error
	if errCheck != nil {
		if errors.Is(errCheck, gorm.ErrRecordNotFound) {
			insertNewItem = true
		} else {
			return nil, errCheck
		}
	}

	if insertNewItem {
		errNewInsert := r.db.Create(&item).Error
		if errNewInsert != nil {
			return nil, errNewInsert
		}
	} else {
		errUpdate := r.db.Model(&item).Where("order_id = ? AND name = ?", item.OrderId, item.Name).Updates(models.Item{
			Name:        item.Name,
			Quantity:    item.Quantity,
			Description: item.Description,
		}).Error
		if errUpdate != nil {
			return nil, errUpdate
		}
	}

	return item, nil

}
