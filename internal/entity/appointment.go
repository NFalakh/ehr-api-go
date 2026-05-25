package entity

import (
	"time"

	"gorm.io/gorm"
)

// Appointment represents a scheduled visit for a patient with a doctor.
type Appointment struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	PatientID       uint           `gorm:"not null;index" json:"patient_id"`
	DoctorID        uint           `gorm:"not null;index" json:"doctor_id"`
	AppointmentDate time.Time      `gorm:"not null" json:"appointment_date"`
	Status          string         `gorm:"type:enum('scheduled','completed','cancelled','no_show');default:'scheduled';not null" json:"status"`
	Reason          string         `gorm:"type:text" json:"reason"`
	Notes           string         `gorm:"type:text" json:"notes"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Patient *Patient `json:"patient,omitempty"`
	Doctor  *User    `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
}
