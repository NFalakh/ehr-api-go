package main

import (
	"log"

	"ehr-api/internal/config"
	"ehr-api/internal/handler"
	"ehr-api/internal/repository"
	"ehr-api/internal/router"
	"ehr-api/internal/usecase"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()

	// Initialize Database
	db, err := cfg.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Database connected and auto-migrated successfully")

	// Dependency Injection for all modules
	// Patient
	patientRepo := repository.NewPatientRepository(db)
	patientUsecase := usecase.NewPatientUsecase(patientRepo)
	patientHandler := handler.NewPatientHandler(patientUsecase)

	// Allergy
	allergyRepo := repository.NewAllergyRepository(db)
	allergyUsecase := usecase.NewAllergyUsecase(allergyRepo)
	allergyHandler := handler.NewAllergyHandler(allergyUsecase)

	// Appointment
	appointmentRepo := repository.NewAppointmentRepository(db)
	appointmentUsecase := usecase.NewAppointmentUsecase(appointmentRepo)
	appointmentHandler := handler.NewAppointmentHandler(appointmentUsecase)

	// Medication
	medicationRepo := repository.NewMedicationRepository(db)
	medicationUsecase := usecase.NewMedicationUsecase(medicationRepo)
	medicationHandler := handler.NewMedicationHandler(medicationUsecase)

	// LabResult
	labResultRepo := repository.NewLabResultRepository(db)
	labResultUsecase := usecase.NewLabResultUsecase(labResultRepo)
	labResultHandler := handler.NewLabResultHandler(labResultUsecase)

	// VitalSign
	vitalSignRepo := repository.NewVitalSignRepository(db)
	vitalSignUsecase := usecase.NewVitalSignUsecase(vitalSignRepo)
	vitalSignHandler := handler.NewVitalSignHandler(vitalSignUsecase)

	// User
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// MedicalRecord
	medicalRecordRepo := repository.NewMedicalRecordRepository(db)
	medicalRecordUsecase := usecase.NewMedicalRecordUsecase(medicalRecordRepo)
	medicalRecordHandler := handler.NewMedicalRecordHandler(medicalRecordUsecase)

	// Setup Router with all handlers
	r := router.SetupRouter(
		patientHandler,
		allergyHandler,
		appointmentHandler,
		medicationHandler,
		labResultHandler,
		vitalSignHandler,
		userHandler,
		medicalRecordHandler,
	)

	// Start Server
	log.Printf("Server starting on port %s...", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
