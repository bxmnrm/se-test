package models
import "gorm.io/gorm"

type StatusNews struct {
	gorm.Model
	StatusNewsID   string `json:"StatusNewsID" gorm:"not null;primaryKey"`
	StatusNewsName string `json:"StatusNewsName" gorm:"not null"`
}