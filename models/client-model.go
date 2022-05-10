package models

//Client model
type Client struct {
	ID      uint   `json:"id"  gorm:"primarykey"`
	Name    string `json:"name"     validate:"required,min=3,max=127"  gorm:"size:127;not null"`
	Surname string `json:"surname"  validate:"required,min=3,max=255"  gorm:"size:255;not null"`
	Email   string `json:"email"    validate:"required,email,max=125"  gorm:"size:125;unique;not nul"`
	Active  bool   `json:"active"   gorm:"default:false"`
}
