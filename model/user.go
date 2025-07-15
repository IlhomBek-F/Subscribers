package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primaryKey;type:uuid" json:"user_id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	StartDate   string `json:"start_date"`
}

func (s *User) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
