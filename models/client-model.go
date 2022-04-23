package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Client model
type Client struct {
	Model   `json:"-" swagger:"ignore"`
	UUID    string `json:"uuid"    validate:"omitempty,uuid"          gorm:"size:37;not null"`
	Name    string `json:"name"    validate:"required,min=2,max=127"  gorm:"size:127;not null"`
	Surname string `json:"surname" validate:"required,min=2,max=255"  gorm:"size:255;not null"`
	Email   string `json:"email"   validate:"required,email"          gorm:"size:125;not nul"`
}

func (client *Client) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("UUID", uuid)
	return nil
}
