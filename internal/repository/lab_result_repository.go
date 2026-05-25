package repository

import (
    "context"
    "ehr-api/internal/entity"
    "gorm.io/gorm"
)

type LabResultRepository interface {
    Create(ctx context.Context, lr *entity.LabResult) error
    FindAll(ctx context.Context) ([]entity.LabResult, error)
    FindByID(ctx context.Context, id uint) (*entity.LabResult, error)
    Update(ctx context.Context, lr *entity.LabResult) error
    Delete(ctx context.Context, id uint) error
}

type labResultRepo struct {
    db *gorm.DB
}

func NewLabResultRepository(db *gorm.DB) LabResultRepository {
    return &labResultRepo{db: db}
}

func (r *labResultRepo) Create(ctx context.Context, lr *entity.LabResult) error {
    return r.db.WithContext(ctx).Create(lr).Error
}

func (r *labResultRepo) FindAll(ctx context.Context) ([]entity.LabResult, error) {
    var list []entity.LabResult
    if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
        return nil, err
    }
    return list, nil
}

func (r *labResultRepo) FindByID(ctx context.Context, id uint) (*entity.LabResult, error) {
    var lr entity.LabResult
    if err := r.db.WithContext(ctx).First(&lr, id).Error; err != nil {
        return nil, err
    }
    return &lr, nil
}

func (r *labResultRepo) Update(ctx context.Context, lr *entity.LabResult) error {
    return r.db.WithContext(ctx).Model(&entity.LabResult{}).Where("id = ?", lr.ID).Updates(lr).Error
}

func (r *labResultRepo) Delete(ctx context.Context, id uint) error {
    return r.db.WithContext(ctx).Delete(&entity.LabResult{}, id).Error
}
