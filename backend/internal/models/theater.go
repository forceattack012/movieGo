package models

type Theater struct {
	Id       int64      `json:"id" autoIncrement:"true"`
	Name     string     `json:"name" validate:"required"`
	IsOpen   bool       `json:"isOpen"`
	ShowTime []ShowTime `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL, ForeignKey:TheaterId;"`
}
