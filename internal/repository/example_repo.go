package repository

import (
	"github.com/myorg/myservice/internal/models"
	"gorm.io/gorm"
)

// ItemRepository defines the abstraction for Item data access
type ItemRepository interface {
	FindAll() ([]models.Item, error)
	FindByID(id uint) (*models.Item, error)
	Create(item *models.Item) error
	Update(item *models.Item) error
	Delete(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

// NewItemRepository injects the db into the repository
func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) FindAll() ([]models.Item, error) {
	var items []models.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *itemRepository) FindByID(id uint) (*models.Item, error) {
	var item models.Item
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Create(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) Update(item *models.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) Delete(id uint) error {
	return r.db.Delete(&models.Item{}, id).Error
}
