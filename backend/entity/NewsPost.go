package models

import "gorm.io/gorm"

type NewsPost struct {
	gorm.Model
	NewsID        string `json:"NewsID" gorm:"not null;primaryKey"`
	Title         string `json:"Title" gorm:"not null"`

	StatusNewsID  string `json:"StatusNewsID" gorm:"not null;foreignKey:StatusNewsID"`
	FilePath      string `json:"FilePath" gorm:"not null"`
	
	ScholarshipID string `json:"ScholarshipID" gorm:"not null;foreignKey:ScholarshipID"`
	//AdminID       string `json:"AdminID" gorm:"not null" gorm:"foreignKey:AdminID"`
}
