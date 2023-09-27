package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Champion struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid"`
	Name    string    `json:"name"`
	City    string    `json:"city"`
	Faccion string    `json:"faccion"`
}

type Champions struct {
	Champions []Champion `json:"champions"`
}

func (champion *Champion) BeforeCreate(tx *gorm.DB) (err error) {
	champion.ID = uuid.New()
	return
}
