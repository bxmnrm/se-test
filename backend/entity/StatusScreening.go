package models
import "gorm.io/gorm"

type StatusScreening struct {
	gorm.Model
	StatusID   string `json:"StatusID" gorm:"not null;primaryKey"`
	StatusName string `json:"StatusName" gorm:"not null"`
}