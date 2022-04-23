package models

import "time"

//Model this is gorm.Model overwrite. I've done it to hide some field that
// do not need to be passad to the client and eather show at swagger model
type Model struct {
	ID        uint       `gorm:"primary_key" json:"-" swagger:"ignore"`
	CreatedAt time.Time  `json:"-" swagger:"ignore"`
	UpdatedAt time.Time  `json:"-" swagger:"ignore"`
	DeletedAt *time.Time `sql:"index" json:"-" swagger:"ignore"`
}
