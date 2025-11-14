package models

import "gorm.io/gorm"

type NewsPost struct {
	gorm.Model
	Title         string `json:"Title" gorm:"not null"`
	StatusNewsID  string `json:"StatusNewsID" gorm:"not null"`
	FilePath      string `json:"FilePath" gorm:"not null"`
	ScholarshipID string `json:"ScholarshipID" gorm:"not null"`
	AdminID       string `json:"AdminID" gorm:"not null"`
}
