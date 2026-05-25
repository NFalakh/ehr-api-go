package repository

import (
	"context"
	"ehr-api/internal/entity"

	"gorm.io/gorm"
)

type PatientRepository interface {
	Create(ctx context.Context, patient *entity.Patient) error
	FindAll(ctx context.Context) ([]entity.Patient, error)
	FindByID(ctx context.Context, id uint) (*entity.Patient, error)
	Update(ctx context.Context, patient *entity.Patient) error
	Delete(ctx context.Context, id uint) error
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) Create(ctx context.Context, patient *entity.Patient) error {
	return r.db.WithContext(ctx).Create(patient).Error
}

func (r *patientRepository) FindAll(ctx context.Context) ([]entity.Patient, error) {
	var patients []entity.Patient
	err := r.db.WithContext(ctx).Find(&patients).Error
	return patients, err
}

func (r *patientRepository) FindByID(ctx context.Context, id uint) (*entity.Patient, error) {
	var patient entity.Patient
	err := r.db.WithContext(ctx).First(&patient, id).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) Update(ctx context.Context, patient *entity.Patient) error {
	return r.db.WithContext(ctx).Save(patient).Error
}

func (r *patientRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Patient{}, id).Error
}
