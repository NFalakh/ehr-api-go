package router

import (
	"ehr-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
    patientHandler *handler.PatientHandler,
    allergyHandler *handler.AllergyHandler,
    appointmentHandler *handler.AppointmentHandler,
    medicationHandler *handler.MedicationHandler,
    labResultHandler *handler.LabResultHandler,
    vitalSignHandler *handler.VitalSignHandler,
    userHandler *handler.UserHandler,
    medicalRecordHandler *handler.MedicalRecordHandler,
) *gin.Engine {
	r := gin.Default()

	// Global middleware
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// API routes
	api := r.Group("/api/v1")
	{
		    // Patient routes
    patients := api.Group("/patients")
    {
        patients.POST("", patientHandler.Create)
        patients.GET("", patientHandler.FindAll)
        patients.GET("/:id", patientHandler.FindByID)
        patients.PUT("/:id", patientHandler.Update)
        patients.DELETE("/:id", patientHandler.Delete)
    }

    // Allergy routes
    allergies := api.Group("/allergies")
    {
        allergies.POST("", allergyHandler.Create)
        allergies.GET("", allergyHandler.FindAll)
        allergies.GET("/:id", allergyHandler.FindByID)
        allergies.PUT("/:id", allergyHandler.Update)
        allergies.DELETE("/:id", allergyHandler.Delete)
    }

    // Appointment routes
    appointments := api.Group("/appointments")
    {
        appointments.POST("", appointmentHandler.Create)
        appointments.GET("", appointmentHandler.FindAll)
        appointments.GET("/:id", appointmentHandler.FindByID)
        appointments.PUT("/:id", appointmentHandler.Update)
        appointments.DELETE("/:id", appointmentHandler.Delete)
    }

    // Medication routes
    medications := api.Group("/medications")
    {
        medications.POST("", medicationHandler.Create)
        medications.GET("", medicationHandler.FindAll)
        medications.GET("/:id", medicationHandler.FindByID)
        medications.PUT("/:id", medicationHandler.Update)
        medications.DELETE("/:id", medicationHandler.Delete)
    }

    // LabResult routes
    labResults := api.Group("/labresults")
    {
        labResults.POST("", labResultHandler.Create)
        labResults.GET("", labResultHandler.FindAll)
        labResults.GET("/:id", labResultHandler.FindByID)
        labResults.PUT("/:id", labResultHandler.Update)
        labResults.DELETE("/:id", labResultHandler.Delete)
    }

    // VitalSign routes
    vitalSigns := api.Group("/vitalsigns")
    {
        vitalSigns.POST("", vitalSignHandler.Create)
        vitalSigns.GET("", vitalSignHandler.FindAll)
        vitalSigns.GET("/:id", vitalSignHandler.FindByID)
        vitalSigns.PUT("/:id", vitalSignHandler.Update)
        vitalSigns.DELETE("/:id", vitalSignHandler.Delete)
    }

    // User routes
    users := api.Group("/users")
    {
        users.POST("", userHandler.Create)
        users.GET("", userHandler.FindAll)
        users.GET("/:id", userHandler.FindByID)
        users.PUT("/:id", userHandler.Update)
        users.DELETE("/:id", userHandler.Delete)
    }

    // MedicalRecord routes
    medicalRecords := api.Group("/medicalrecords")
    {
        medicalRecords.POST("", medicalRecordHandler.Create)
        medicalRecords.GET("", medicalRecordHandler.FindAll)
        medicalRecords.GET("/:id", medicalRecordHandler.FindByID)
        medicalRecords.PUT("/:id", medicalRecordHandler.Update)
        medicalRecords.DELETE("/:id", medicalRecordHandler.Delete)
    }
	}

	return r
}
