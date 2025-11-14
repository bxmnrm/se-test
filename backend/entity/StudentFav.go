package models
import "gorm.io/gorm"

type StudentFav struct {
	gorm.Model
	FavID     string `json:"FavID" gorm:"not null;primaryKey"`
	StudentID string `json:"StudentID" gorm:"foreignKey:StudentID"`
	NewsID    string `json:"NewsID" gorm:"foreignKey:NewsID"`
}