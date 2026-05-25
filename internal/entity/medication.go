package entity

import (
	"time"

	"gorm.io/gorm"
)

// Medication represents a drug prescribed to a patient during a medical record encounter.
type Medication struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	MedicalRecordID uint           `gorm:"not null;index" json:"medical_record_id"`
	Name            string         `gorm:"type:varchar(100);not null" json:"name"`
	Dosage          string         `gorm:"type:varchar(50);not null" json:"dosage"`
	Frequency       string         `gorm:"type:varchar(50);not null" json:"frequency"`
	Duration        string         `gorm:"type:varchar(50)" json:"duration"` // e.g., "7 days"
	Notes           string         `gorm:"type:text" json:"notes"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
