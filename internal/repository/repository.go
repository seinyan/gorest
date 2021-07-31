package repository

import "gorm.io/gorm"

type repository struct {
	conn *gorm.DB
}