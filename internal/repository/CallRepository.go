package repository

import (
	"github.com/seinyan/gorest/internal/models"
	"gorm.io/gorm"
)

type CallRepository interface {
	CreateItems(items []models.Call) ([]models.Call, error)
}

type callRepository struct {
	conn *gorm.DB
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