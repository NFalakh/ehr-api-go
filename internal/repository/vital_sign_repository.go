package repository

import (
    "context"
    "ehr-api/internal/entity"
    "gorm.io/gorm"
)

type VitalSignRepository interface {
    Create(ctx context.Context, vs *entity.VitalSign) error
    FindAll(ctx context.Context) ([]entity.VitalSign, error)
    FindByID(ctx context.Context, id uint) (*entity.VitalSign, error)
    Update(ctx context.Context, vs *entity.VitalSign) error
    Delete(ctx context.Context, id uint) error
}

type vitalSignRepo struct {
    db *gorm.DB
}

func NewVitalSignRepository(db *gorm.DB) VitalSignRepository {
    return &vitalSignRepo{db: db}
}

func (r *vitalSignRepo) Create(ctx context.Context, vs *entity.VitalSign) error {
    return r.db.WithContext(ctx).Create(vs).Error
}

func (r *vitalSignRepo) FindAll(ctx context.Context) ([]entity.VitalSign, error) {
    var list []entity.VitalSign
    if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
        return nil, err
    }
    return list, nil
}

func (r *vitalSignRepo) FindByID(ctx context.Context, id uint) (*entity.VitalSign, error) {
    var vs entity.VitalSign
    if err := r.db.WithContext(ctx).First(&vs, id).Error; err != nil {
        return nil, err
    }
    return &vs, nil
}

func (r *vitalSignRepo) Update(ctx context.Context, vs *entity.VitalSign) error {
    return r.db.WithContext(ctx).Model(&entity.VitalSign{}).Where("id = ?", vs.ID).Updates(vs).Error
}

func (r *vitalSignRepo) Delete(ctx context.Context, id uint) error {
    return r.db.WithContext(ctx).Delete(&entity.VitalSign{}, id).Error
}
