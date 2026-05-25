package repository

import (
    "context"
    "ehr-api/internal/entity"
    "gorm.io/gorm"
)

type MedicalRecordRepository interface {
    Create(ctx context.Context, mr *entity.MedicalRecord) error
    FindAll(ctx context.Context) ([]entity.MedicalRecord, error)
    FindByID(ctx context.Context, id uint) (*entity.MedicalRecord, error)
    Update(ctx context.Context, mr *entity.MedicalRecord) error
    Delete(ctx context.Context, id uint) error
}

type medicalRecordRepo struct {
    db *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) MedicalRecordRepository {
    return &medicalRecordRepo{db: db}
}

func (r *medicalRecordRepo) Create(ctx context.Context, mr *entity.MedicalRecord) error {
    return r.db.WithContext(ctx).Create(mr).Error
}

func (r *medicalRecordRepo) FindAll(ctx context.Context) ([]entity.MedicalRecord, error) {
    var list []entity.MedicalRecord
    if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
        return nil, err
    }
    return list, nil
}

func (r *medicalRecordRepo) FindByID(ctx context.Context, id uint) (*entity.MedicalRecord, error) {
    var mr entity.MedicalRecord
    if err := r.db.WithContext(ctx).First(&mr, id).Error; err != nil {
        return nil, err
    }
    return &mr, nil
}

func (r *medicalRecordRepo) Update(ctx context.Context, mr *entity.MedicalRecord) error {
    return r.db.WithContext(ctx).Model(&entity.MedicalRecord{}).Where("id = ?", mr.ID).Updates(mr).Error
}

func (r *medicalRecordRepo) Delete(ctx context.Context, id uint) error {
    return r.db.WithContext(ctx).Delete(&entity.MedicalRecord{}, id).Error
}
