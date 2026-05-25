package entity

import (
	"time"

	"gorm.io/gorm"
)

// User represents a staff member (doctor, nurse, admin) in the system.
type User struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"type:varchar(100);not null" json:"name"`
	Username       string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Password       string         `gorm:"type:varchar(255);not null" json:"-"`
	Role           string         `gorm:"type:enum('doctor','nurse','admin');not null" json:"role"`
	Specialization string         `gorm:"type:varchar(100)" json:"specialization,omitempty"`
	Phone          string         `gorm:"type:varchar(20)" json:"phone,omitempty"`
	Email          string         `gorm:"type:varchar(100);uniqueIndex" json:"email,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
