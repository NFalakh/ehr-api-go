package repository

import (
    "context"
    "ehr-api/internal/entity"
    "gorm.io/gorm"
)

type AllergyRepository interface {
    Create(ctx context.Context, allergy *entity.Allergy) error
    FindAll(ctx context.Context) ([]entity.Allergy, error)
    FindByID(ctx context.Context, id uint) (*entity.Allergy, error)
    Update(ctx context.Context, allergy *entity.Allergy) error
    Delete(ctx context.Context, id uint) error
}

type allergyRepo struct {
    db *gorm.DB
}

func NewAllergyRepository(db *gorm.DB) AllergyRepository {
    return &allergyRepo{db: db}
}

func (r *allergyRepo) Create(ctx context.Context, allergy *entity.Allergy) error {
    return r.db.WithContext(ctx).Create(allergy).Error
}

func (r *allergyRepo) FindAll(ctx context.Context) ([]entity.Allergy, error) {
    var allergies []entity.Allergy
    if err := r.db.WithContext(ctx).Find(&allergies).Error; err != nil {
        return nil, err
    }
    return allergies, nil
}

func (r *allergyRepo) FindByID(ctx context.Context, id uint) (*entity.Allergy, error) {
    var allergy entity.Allergy
    if err := r.db.WithContext(ctx).First(&allergy, id).Error; err != nil {
        return nil, err
    }
    return &allergy, nil
}

func (r *allergyRepo) Update(ctx context.Context, allergy *entity.Allergy) error {
    return r.db.WithContext(ctx).Model(&entity.Allergy{}).Where("id = ?", allergy.ID).Updates(allergy).Error
}

func (r *allergyRepo) Delete(ctx context.Context, id uint) error {
    return r.db.WithContext(ctx).Delete(&entity.Allergy{}, id).Error
}
