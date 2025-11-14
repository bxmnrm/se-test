package models

import (
	"gorm.io/gorm"
)

type Industry struct {
	gorm.Model
	Name		string		`json:"name"`
}