package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	ID     uint    `json:"id" gorm:"primary_key"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
