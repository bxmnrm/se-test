package models

import "gorm.io/gorm"

type Screening struct {
    gorm.Model

    //CriterialID   string `json:"CriterialID"`
    //Criteria      Criteria `gorm:"foreignKey:CriterialID"`

    //ApplicationID string `json:"ApplicationID"`
    //Application   Application `gorm:"foreignKey:ApplicationID"`

    StatusID      string `json:"StatusID"`
    Status        StatusScreening `gorm:"foreignKey:StatusID"`

    RejectReason  string `json:"RejectReason"`
}