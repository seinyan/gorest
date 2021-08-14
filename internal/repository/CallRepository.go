package repository

import (
	"errors"
	"github.com/seinyan/gorest/internal/models"
	"gorm.io/gorm"
)

type CallRepository interface {
	CreateItems(items []models.Call) ([]models.Call, error)
	GetItems() ([]models.Call, error)
}

type callRepository struct {
	conn *gorm.DB
}

func (repo callRepository) GetItems() ([]models.Call, error) {
	return nil, errors.New("Err")
}

func NewCallRepository(conn *gorm.DB) CallRepository {
	return &callRepository{
		conn: conn,
	}
}


func (repo callRepository) CreateItems(items []models.Call) ([]models.Call, error) {
	err := repo.conn.Create(&items).Error
	return items, err
}