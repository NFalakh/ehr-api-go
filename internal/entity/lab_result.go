package entity

import (
	"time"

	"gorm.io/gorm"
)

// LabResult represents laboratory test results for a patient.
type LabResult struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	MedicalRecordID uint           `gorm:"not null;index" json:"medical_record_id"`
	TestName        string         `gorm:"type:varchar(100);not null" json:"test_name"`
	Result          string         `gorm:"type:varchar(255);not null" json:"result"`
	Unit            string         `gorm:"type:varchar(50)" json:"unit"`
	ReferenceRange  string         `gorm:"type:varchar(100)" json:"reference_range"`
	TestDate        time.Time      `gorm:"not null" json:"test_date"`
	Remarks         string         `gorm:"type:text" json:"remarks"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
