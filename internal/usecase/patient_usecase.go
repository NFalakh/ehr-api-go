package usecase

import (
	"context"
	"ehr-api/internal/dto"
	"ehr-api/internal/entity"
	"ehr-api/internal/repository"
	"time"
)

type PatientUsecase interface {
	Create(ctx context.Context, req dto.CreatePatientRequest) (*dto.PatientResponse, error)
	FindAll(ctx context.Context) ([]dto.PatientResponse, error)
	FindByID(ctx context.Context, id uint) (*dto.PatientResponse, error)
	Update(ctx context.Context, id uint, req dto.UpdatePatientRequest) (*dto.PatientResponse, error)
	Delete(ctx context.Context, id uint) error
}

type patientUsecase struct {
	repo repository.PatientRepository
}

func NewPatientUsecase(repo repository.PatientRepository) PatientUsecase {
	return &patientUsecase{repo: repo}
}

func (u *patientUsecase) Create(ctx context.Context, req dto.CreatePatientRequest) (*dto.PatientResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, err
	}

	patient := &entity.Patient{
		IdentityNumber: req.IdentityNumber,
		Name:           req.Name,
		DateOfBirth:    dob,
		Gender:         req.Gender,
		Address:        req.Address,
		Phone:          req.Phone,
		Email:          req.Email,
		BloodType:      req.BloodType,
	}

	if err := u.repo.Create(ctx, patient); err != nil {
		return nil, err
	}

	return toPatientResponse(patient), nil
}

func (u *patientUsecase) FindAll(ctx context.Context) ([]dto.PatientResponse, error) {
	patients, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.PatientResponse
	for _, p := range patients {
		responses = append(responses, *toPatientResponse(&p))
	}

	return responses, nil
}

func (u *patientUsecase) FindByID(ctx context.Context, id uint) (*dto.PatientResponse, error) {
	patient, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toPatientResponse(patient), nil
}

func (u *patientUsecase) Update(ctx context.Context, id uint, req dto.UpdatePatientRequest) (*dto.PatientResponse, error) {
	patient, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		patient.Name = req.Name
	}
	if req.Gender != "" {
		patient.Gender = req.Gender
	}
	if req.Address != "" {
		patient.Address = req.Address
	}
	if req.Phone != "" {
		patient.Phone = req.Phone
	}
	if req.Email != "" {
		patient.Email = req.Email
	}
	if req.BloodType != "" {
		patient.BloodType = req.BloodType
	}

	if err := u.repo.Update(ctx, patient); err != nil {
		return nil, err
	}

	return toPatientResponse(patient), nil
}

func (u *patientUsecase) Delete(ctx context.Context, id uint) error {
	return u.repo.Delete(ctx, id)
}

func toPatientResponse(patient *entity.Patient) *dto.PatientResponse {
	return &dto.PatientResponse{
		ID:             patient.ID,
		IdentityNumber: patient.IdentityNumber,
		Name:           patient.Name,
		DateOfBirth:    patient.DateOfBirth,
		Gender:         patient.Gender,
		Address:        patient.Address,
		Phone:          patient.Phone,
		Email:          patient.Email,
		BloodType:      patient.BloodType,
		CreatedAt:      patient.CreatedAt,
		UpdatedAt:      patient.UpdatedAt,
	}
}
