package repository

import (
    "context"
    "ehr-api/internal/entity"
    "gorm.io/gorm"
)

type MedicationRepository interface {
    Create(ctx context.Context, med *entity.Medication) error
    FindAll(ctx context.Context) ([]entity.Medication, error)
    FindByID(ctx context.Context, id uint) (*entity.Medication, error)
    Update(ctx context.Context, med *entity.Medication) error
    Delete(ctx context.Context, id uint) error
}

type medicationRepo struct {
    db *gorm.DB
}

func NewMedicationRepository(db *gorm.DB) MedicationRepository {
    return &medicationRepo{db: db}
}

func (r *medicationRepo) Create(ctx context.Context, med *entity.Medication) error {
    return r.db.WithContext(ctx).Create(med).Error
}

func (r *medicationRepo) FindAll(ctx context.Context) ([]entity.Medication, error) {
    var meds []entity.Medication
    if err := r.db.WithContext(ctx).Find(&meds).Error; err != nil {
        return nil, err
    }
    return meds, nil
}

func (r *medicationRepo) FindByID(ctx context.Context, id uint) (*entity.Medication, error) {
    var med entity.Medication
    if err := r.db.WithContext(ctx).First(&med, id).Error; err != nil {
        return nil, err
    }
    return &med, nil
}

func (r *medicationRepo) Update(ctx context.Context, med *entity.Medication) error {
    return r.db.WithContext(ctx).Model(&entity.Medication{}).Where("id = ?", med.ID).Updates(med).Error
}

func (r *medicationRepo) Delete(ctx context.Context, id uint) error {
    return r.db.WithContext(ctx).Delete(&entity.Medication{}, id).Error
}
