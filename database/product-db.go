package database

import (
	"gorm.io/gorm"
)

type ProductDB struct {
	gorm.Model
	Title string 
	Description  string
}