package entity

import (
	"time"

	"gorm.io/gorm"
)

// Allergy represents a patient's allergy.
type Allergy struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	PatientID   uint           `gorm:"not null;index" json:"patient_id"`
	Allergen    string         `gorm:"type:varchar(100);not null" json:"allergen"`
	Severity    string         `gorm:"type:enum('mild','moderate','severe');not null" json:"severity"`
	Reaction    string         `gorm:"type:varchar(255)" json:"reaction"`
	NotedDate   time.Time      `gorm:"type:date" json:"noted_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
