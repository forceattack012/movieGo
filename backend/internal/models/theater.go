package models

type Theater struct {
	Id         int64  `json:"Id" validate:"required"`
	Name       string `json:"Name" validate:"required"`
	ScreenName string `json:"ScreenName" validate:"required"`
	IsOpen     bool   `json:"IsOpen" validate:"required"`
}
