package models

import (
	"gorm.io/gorm"
	"time"
)

type Call struct {
	gorm.Model
	ID          uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	At          time.Time
	Phone       string                `gorm:"type:varchar;index"`
	Description string                `gorm:"type:text;" json:"description"`
	Histories   []CallHistory `gorm:"foreignKey:CallId"`
}

type CallHistory struct {
	ID     uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	At     time.Time
	CallId uint64 `gorm:"type:bigint;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"call_id"`
	Value  string `gorm:"type:text;" json:"value"`
}

func (u Call) TableName() string {
	return "calls"
}

func (u CallHistory) TableName() string {
	return "calls_history"
}
