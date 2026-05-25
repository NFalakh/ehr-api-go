package entity

import (
	"time"

	"gorm.io/gorm"
)

// MedicalRecord represents a single encounter/visit of a patient.
type MedicalRecord struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	PatientID        uint           `gorm:"not null;index" json:"patient_id"`
	DoctorID         uint           `gorm:"not null;index" json:"doctor_id"`
	VisitDate        time.Time      `gorm:"not null" json:"visit_date"`
	ChiefComplaint   string         `gorm:"type:text;not null" json:"chief_complaint"`
	Diagnosis        string         `gorm:"type:text;not null" json:"diagnosis"`
	TreatmentPlan    string         `gorm:"type:text" json:"treatment_plan"`
	Treatment        string         `gorm:"type:varchar(255)" json:"treatment"`
	Prescription     string         `gorm:"type:text" json:"prescription"`
	Notes            string         `gorm:"type:text" json:"notes"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Patient     *Patient     `json:"patient,omitempty"`
	Doctor      *User        `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
	Medications []Medication `json:"medications,omitempty"`
	VitalSigns  []VitalSign  `json:"vital_signs,omitempty"`
	LabResults  []LabResult  `json:"lab_results,omitempty"`
}
