package models

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

//Client model
type Client struct {
	Model   `json:"-" swagger:"ignore"`
	UUID    uuid.UUID `json:"uuid"    validate:"omitempty,uuid"          gorm:"type:uuid;default:uuid_generate_v4()"`
	Name    string    `json:"name"    validate:"required,min=2,max=127"  gorm:"size:127;not null"`
	Surname string    `json:"surname" validate:"required,min=2,max=255"  gorm:"size: 255;not null"`
	Email   string    `json:"email"   validate:"required,email"          gorm:"size:125;not nul"`
}

//CreateNewClient function to create new client. This function validate the model after create it. In case of validation
//failure, this function will return a tuple (nil, error), other else (Client, nil)
func CreateNewClient(surname string, name string, email string) (*Client, error) {
	client := Client{
		UUID:    uuid.New(),
		Surname: surname,
		Name:    name,
		Email:   email,
	}

	modelValidate := validator.New()
	err := modelValidate.Struct(client)

	if err != nil {
		return nil, err
	}

	return &client, nil
}
