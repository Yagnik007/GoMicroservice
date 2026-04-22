package services

import (
	"github.com/myorg/myservice/internal/models"
	"github.com/myorg/myservice/internal/repository"
)

// ItemService defines business logic abstraction for Items
type ItemService interface {
	GetAllItems() ([]models.Item, error)
	GetItemByID(id uint) (*models.Item, error)
	CreateItem(item *models.Item) error
}

type itemService struct {
	repo repository.ItemRepository
}

// NewItemService injects the repository dependency
func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{repo}
}

func (s *itemService) GetAllItems() ([]models.Item, error) {
	return s.repo.FindAll()
}

func (s *itemService) GetItemByID(id uint) (*models.Item, error) {
	return s.repo.FindByID(id)
}

func (s *itemService) CreateItem(item *models.Item) error {
	// Add potential business logic or validation here
	return s.repo.Create(item)
}
