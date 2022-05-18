package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Client model
type Client struct {
	ID      string  `json:"id"  gorm:"primarykey;->;<-:create"`
	Name    *string `json:"name"     validate:"required,min=3"  gorm:"not null"`
	Surname *string `json:"surname"  validate:"required,min=3"  gorm:"not null"`
	Email   *string `json:"email"    validate:"required,email"  gorm:"not null;unique;index"`
	Active  bool    `json:"active"   gorm:"default:false"`
}

func NewClient(name, surname, email string) *Client {
	return &Client{
		Name:    GetStringPointer(name),
		Surname: GetStringPointer(surname),
		Email:   GetStringPointer(email),
	}
}

func (client *Client) BeforeCreate(tx *gorm.DB) (err error) {
	client.ID = uuid.NewString()

	client.ToLower()

	return
}

func (client *Client) BeforeUpdate(tx *gorm.DB) (err error) {
	client.ToLower()

	return
}

func (client *Client) ToLower() {
	toLower(client.Name)
	toLower(client.Surname)
	toLower(client.Email)
}

func (client *Client) Equals(other *Client) bool {
	if other == nil {
		return false
	}

	if client.ID != other.ID {
		return false
	}

	if *client.Name != *other.Name {
		return false
	}

	if *client.Surname != *other.Surname {
		return false
	}

	if *client.Email != *other.Email {
		return false
	}

	if client.Active != other.Active {
		return false
	}

	return true
}

func (client *Client) ToString() string {
	return fmt.Sprintf(
		"[%s %s %s %s %v",
		client.ID,
		*client.Name,
		*client.Surname,
		*client.Email,
		client.Active,
	)
}
