package models

import (
	"gorm.io/gorm"
)

// Initiate Models
type LoanDetailsLedger struct {
	ID              uint `gorm:"primaryKey"`
	BankName        string
	BorrowerName    string
	PrincipalAmount float64
	NoOfYears       int
	RateOfInterest  float64
	EmiId           string
	Closed          int `gorm:"default:0"`
}

type EmiPaymentDetailLedger struct {
	ID            uint `gorm:"primaryKey"`
	EmiId         string
	EmiAmount     float64
	EmiPaidAmount float64
	EmiYear       int
	EmiMonth      int
	EmiDay        int
	PaymentStatus int `gorm:"default:0"`
	AddedLumpsum  int `gorm:"default:0"`
}

func AutoMigrateModel(DBC *gorm.DB) {
	DBC.AutoMigrate(&LoanDetailsLedger{})
	DBC.AutoMigrate(&EmiPaymentDetailLedger{})
}
