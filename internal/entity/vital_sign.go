package entity

import (
	"time"

	"gorm.io/gorm"
)

// VitalSign represents patient vital signs taken during an encounter.
type VitalSign struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	MedicalRecordID uint           `gorm:"not null;index" json:"medical_record_id"`
	BloodPressure   string         `gorm:"type:varchar(20)" json:"blood_pressure"` // e.g., 120/80
	HeartRate       int            `json:"heart_rate"`                             // bpm
	Temperature     float64        `json:"temperature"`                            // Celsius
	RespiratoryRate int            `json:"respiratory_rate"`                       // breaths per minute
	OxygenSaturation int           `json:"oxygen_saturation"`                      // percentage
	Weight          float64        `json:"weight"`                                 // kg
	Height          float64        `json:"height"`                                 // cm
	RecordedAt      time.Time      `gorm:"not null" json:"recorded_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
