package repository

import (
    "context"
    "ehr-api/internal/entity"
    "gorm.io/gorm"
)

type AppointmentRepository interface {
    Create(ctx context.Context, appt *entity.Appointment) error
    FindAll(ctx context.Context) ([]entity.Appointment, error)
    FindByID(ctx context.Context, id uint) (*entity.Appointment, error)
    Update(ctx context.Context, appt *entity.Appointment) error
    Delete(ctx context.Context, id uint) error
}

type appointmentRepo struct {
    db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
    return &appointmentRepo{db: db}
}

func (r *appointmentRepo) Create(ctx context.Context, appt *entity.Appointment) error {
    return r.db.WithContext(ctx).Create(appt).Error
}

func (r *appointmentRepo) FindAll(ctx context.Context) ([]entity.Appointment, error) {
    var appts []entity.Appointment
    if err := r.db.WithContext(ctx).Find(&appts).Error; err != nil {
        return nil, err
    }
    return appts, nil
}

func (r *appointmentRepo) FindByID(ctx context.Context, id uint) (*entity.Appointment, error) {
    var appt entity.Appointment
    if err := r.db.WithContext(ctx).First(&appt, id).Error; err != nil {
        return nil, err
    }
    return &appt, nil
}

func (r *appointmentRepo) Update(ctx context.Context, appt *entity.Appointment) error {
    return r.db.WithContext(ctx).Model(&entity.Appointment{}).Where("id = ?", appt.ID).Updates(appt).Error
}

func (r *appointmentRepo) Delete(ctx context.Context, id uint) error {
    return r.db.WithContext(ctx).Delete(&entity.Appointment{}, id).Error
}
