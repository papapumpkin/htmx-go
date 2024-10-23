package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	PatientFirstName  string `gorm:"column:patient_first_name;not null"`
	PatientLastName   string `gorm:"column:patient_last_name;not null"`
	ProviderFirstName string `gorm:"column:provider_first_name;not null"`
	ProviderLastName  string `gorm:"column:provider_last_name;not null"`
}

type OrderRequest struct {
	PatientFirstName  string `form:"patient_first_name" json:"patient_first_name" binding:"required"`
	PatientLastName   string `form:"patient_last_name" json:"patient_last_name" binding:"required"`
	ProviderFirstName string `form:"provider_first_name" json:"provider_first_name" binding:"required"`
	ProviderLastName  string `form:"provider_last_name" json:"provider_last_name" binding:"required"`
}
