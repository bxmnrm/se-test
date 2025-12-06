package models
import "gorm.io/gorm"

type StudentFav struct {
    gorm.Model

    FavID     string `json:"FavID" gorm:"not null;primaryKey"`

    //StudentID string `json:"StudentID" gorm:"not null"`
    //Student   Student `gorm:"foreignKey:StudentID"`

    NewsID    string `json:"NewsID" gorm:"not null"`
    NewsPost  NewsPost `gorm:"foreignKey:NewsID"`
}
