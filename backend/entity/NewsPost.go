package models

import "gorm.io/gorm"

type NewsPost struct {
	gorm.Model

	NewsID 	  string `json:"NewsID" gorm:"not null;primaryKey"`
	Title     string `json:"Title" gorm:"not null"`

	StatusNewsID string `json:"StatusNewsID" gorm:"not null"`
	StatusNews   StatusNews `gorm:"foreignKey:StatusNewsID"`

	FilePath  string `json:"FilePath" gorm:"not null"`

	//ScholarshipID string `json:"ScholarshipID" gorm:"not null"`
	//Scholarship   Scholarship `gorm:"foreignKey:ScholarshipID"`

	//AdminID   string `json:"AdminID" gorm:"not null"`
	//Admin     Admin `gorm:"foreignKey:AdminID"`
}
