package repository

import (
    "context"
    "ehr-api/internal/entity"
    "gorm.io/gorm"
)

type UserRepository interface {
    Create(ctx context.Context, user *entity.User) error
    FindAll(ctx context.Context) ([]entity.User, error)
    FindByID(ctx context.Context, id uint) (*entity.User, error)
    Update(ctx context.Context, user *entity.User) error
    Delete(ctx context.Context, id uint) error
}

type userRepo struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *entity.User) error {
    return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) FindAll(ctx context.Context) ([]entity.User, error) {
    var users []entity.User
    if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (r *userRepo) FindByID(ctx context.Context, id uint) (*entity.User, error) {
    var user entity.User
    if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepo) Update(ctx context.Context, user *entity.User) error {
    return r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *userRepo) Delete(ctx context.Context, id uint) error {
    return r.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}
