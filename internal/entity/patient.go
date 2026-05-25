package entity

import (
	"time"

	"gorm.io/gorm"
)

// Patient represents a patient in the EHR system.
type Patient struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	IdentityNumber string      `gorm:"type:varchar(50);uniqueIndex;not null" json:"identity_number"` // e.g., NIK/KTP
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	DateOfBirth time.Time      `gorm:"type:date;not null" json:"date_of_birth"`
	Gender      string         `gorm:"type:enum('male','female','other');not null" json:"gender"`
	Address     string         `gorm:"type:text" json:"address"`
	Phone       string         `gorm:"type:varchar(20)" json:"phone"`
	Email       string         `gorm:"type:varchar(100)" json:"email,omitempty"`
	BloodType   string         `gorm:"type:varchar(5)" json:"blood_type,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	MedicalRecords []MedicalRecord `json:"medical_records,omitempty"`
	Allergies      []Allergy       `json:"allergies,omitempty"`
	Appointments   []Appointment   `json:"appointments,omitempty"`
}
