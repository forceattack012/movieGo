package models

import "gorm.io/gorm"

type Theater struct {
	gorm.Model
	Id       int64      `json:"id" autoIncrement:"true"`
	Name     string     `json:"name" validate:"required"`
	IsOpen   bool       `json:"isOpen"`
	Size     string     `json:"size" validate:"required"`
	ShowTime []ShowTime `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL, ForeignKey:TheaterId;"`
	Seat     []Seat     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL, ForeignKey:TheaterId;"`
}
