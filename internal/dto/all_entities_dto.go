package dto

// --- Patient DTOs ---

type CreatePatientRequest struct {
    IdentityNumber string `json:"identity_number" binding:"required"`
    Name           string `json:"name" binding:"required"`
    DateOfBirth    string `json:"date_of_birth" binding:"required"` // format: 2006-01-02
    Gender         string `json:"gender" binding:"required,oneof=male female other"`
    BloodType      string `json:"blood_type"`
    Address        string `json:"address"`
    Phone          string `json:"phone"`
    Email          string `json:"email"`
}

type UpdatePatientRequest struct {
    Name        string `json:"name"`
    DateOfBirth string `json:"date_of_birth"`
    Gender      string `json:"gender" binding:"omitempty,oneof=male female other"`
    BloodType   string `json:"blood_type"`
    Address     string `json:"address"`
    Phone       string `json:"phone"`
    Email       string `json:"email"`
}

// --- Allergy DTOs ---

type CreateAllergyRequest struct {
    PatientID uint   `json:"patient_id" binding:"required"`
    Allergen  string `json:"allergen" binding:"required"`
    Severity  string `json:"severity" binding:"required,oneof=mild moderate severe"`
    Reaction  string `json:"reaction"`
    NotedDate string `json:"noted_date"` // ISO date string, parsed in usecase
}

type UpdateAllergyRequest struct {
    Allergen string `json:"allergen"`
    Severity string `json:"severity" binding:"omitempty,oneof=mild moderate severe"`
    Reaction string `json:"reaction"`
    NotedDate string `json:"noted_date"`
}

type AllergyResponse struct {
    ID        uint   `json:"id"`
    PatientID uint   `json:"patient_id"`
    Allergen  string `json:"allergen"`
    Severity  string `json:"severity"`
    Reaction  string `json:"reaction"`
    NotedDate string `json:"noted_date"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

// --- Appointment DTOs ---

type CreateAppointmentRequest struct {
    PatientID       uint   `json:"patient_id" binding:"required"`
    DoctorID        uint   `json:"doctor_id" binding:"required"`
    AppointmentDate string `json:"appointment_date" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
    Status          string `json:"status" binding:"omitempty,oneof=scheduled completed cancelled no_show"`
    Reason          string `json:"reason"`
    Notes           string `json:"notes"`
}

type UpdateAppointmentRequest struct {
    AppointmentDate string `json:"appointment_date" binding:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
    Status          string `json:"status" binding:"omitempty,oneof=scheduled completed cancelled no_show"`
    Reason          string `json:"reason"`
    Notes           string `json:"notes"`
}

type AppointmentResponse struct {
    ID              uint   `json:"id"`
    PatientID       uint   `json:"patient_id"`
    DoctorID        uint   `json:"doctor_id"`
    AppointmentDate string `json:"appointment_date"`
    Status          string `json:"status"`
    Reason          string `json:"reason"`
    Notes           string `json:"notes"`
    CreatedAt       string `json:"created_at"`
    UpdatedAt       string `json:"updated_at"`
}

// --- Medication DTOs ---

type CreateMedicationRequest struct {
    MedicalRecordID uint   `json:"medical_record_id" binding:"required"`
    Name            string `json:"name" binding:"required"`
    Dosage          string `json:"dosage" binding:"required"`
    Frequency       string `json:"frequency" binding:"required"`
    Duration        string `json:"duration"`
    Notes           string `json:"notes"`
}

type UpdateMedicationRequest struct {
    Name      string `json:"name"`
    Dosage    string `json:"dosage"`
    Frequency string `json:"frequency"`
    Duration  string `json:"duration"`
    Notes     string `json:"notes"`
}

type MedicationResponse struct {
    ID              uint   `json:"id"`
    MedicalRecordID uint   `json:"medical_record_id"`
    Name            string `json:"name"`
    Dosage          string `json:"dosage"`
    Frequency       string `json:"frequency"`
    Duration        string `json:"duration"`
    Notes           string `json:"notes"`
    CreatedAt       string `json:"created_at"`
    UpdatedAt       string `json:"updated_at"`
}

// --- LabResult DTOs ---

type CreateLabResultRequest struct {
    MedicalRecordID uint   `json:"medical_record_id" binding:"required"`
    TestName        string `json:"test_name" binding:"required"`
    Result          string `json:"result" binding:"required"`
    Unit            string `json:"unit"`
    ReferenceRange  string `json:"reference_range"`
    TestDate        string `json:"test_date" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
    Remarks         string `json:"remarks"`
}

type UpdateLabResultRequest struct {
    TestName       string `json:"test_name"`
    Result         string `json:"result"`
    Unit           string `json:"unit"`
    ReferenceRange string `json:"reference_range"`
    TestDate       string `json:"test_date" binding:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
    Remarks        string `json:"remarks"`
}

type LabResultResponse struct {
    ID              uint   `json:"id"`
    MedicalRecordID uint   `json:"medical_record_id"`
    TestName        string `json:"test_name"`
    Result          string `json:"result"`
    Unit            string `json:"unit"`
    ReferenceRange  string `json:"reference_range"`
    TestDate        string `json:"test_date"`
    Remarks         string `json:"remarks"`
    CreatedAt       string `json:"created_at"`
    UpdatedAt       string `json:"updated_at"`
}

// --- VitalSign DTOs ---

type CreateVitalSignRequest struct {
    MedicalRecordID uint    `json:"medical_record_id" binding:"required"`
    BloodPressure   string  `json:"blood_pressure"`
    HeartRate       int     `json:"heart_rate"`
    Temperature     float64 `json:"temperature"`
    RespiratoryRate int     `json:"respiratory_rate"`
    OxygenSaturation int    `json:"oxygen_saturation"`
    Weight          float64 `json:"weight"`
    Height          float64 `json:"height"`
    RecordedAt      string   `json:"recorded_at" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

type UpdateVitalSignRequest struct {
    BloodPressure    string  `json:"blood_pressure"`
    HeartRate        *int    `json:"heart_rate"`
    Temperature      *float64 `json:"temperature"`
    RespiratoryRate  *int    `json:"respiratory_rate"`
    OxygenSaturation *int    `json:"oxygen_saturation"`
    Weight           *float64 `json:"weight"`
    Height           *float64 `json:"height"`
    RecordedAt       string   `json:"recorded_at" binding:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}

type VitalSignResponse struct {
    ID              uint    `json:"id"`
    MedicalRecordID uint    `json:"medical_record_id"`
    BloodPressure   string  `json:"blood_pressure"`
    HeartRate       int     `json:"heart_rate"`
    Temperature     float64 `json:"temperature"`
    RespiratoryRate int    `json:"respiratory_rate"`
    OxygenSaturation int   `json:"oxygen_saturation"`
    Weight          float64 `json:"weight"`
    Height          float64 `json:"height"`
    RecordedAt      string  `json:"recorded_at"`
    CreatedAt       string  `json:"created_at"`
    UpdatedAt       string  `json:"updated_at"`
}

// --- User DTOs ---

type CreateUserRequest struct {
    Name           string `json:"name" binding:"required"`
    Username       string `json:"username" binding:"required"`
    Password       string `json:"password" binding:"required"`
    Role           string `json:"role" binding:"required,oneof=doctor nurse admin"`
    Specialization string `json:"specialization"`
    Phone          string `json:"phone"`
    Email          string `json:"email"`
}

type UpdateUserRequest struct {
    Name           *string `json:"name"`
    Username       *string `json:"username"`
    Password       *string `json:"password"`
    Role           *string `json:"role" binding:"omitempty,oneof=doctor nurse admin"`
    Specialization *string `json:"specialization"`
    Phone          *string `json:"phone"`
    Email          *string `json:"email"`
}

type UserResponse struct {
    ID             uint   `json:"id"`
    Name           string `json:"name"`
    Username       string `json:"username"`
    Role           string `json:"role"`
    Specialization string `json:"specialization,omitempty"`
    Phone          string `json:"phone,omitempty"`
    Email          string `json:"email,omitempty"`
    CreatedAt      string `json:"created_at"`
    UpdatedAt      string `json:"updated_at"`
}

// --- MedicalRecord DTOs ---

type CreateMedicalRecordRequest struct {
    PatientID      uint   `json:"patient_id" binding:"required"`
    DoctorID       uint   `json:"doctor_id" binding:"required"`
    VisitDate      string `json:"visit_date" binding:"required,datetime=2006-01-02"`
    ChiefComplaint string `json:"chief_complaint" binding:"required"`
    Diagnosis      string `json:"diagnosis" binding:"required"`
    Treatment      string `json:"treatment"`
    TreatmentPlan  string `json:"treatment_plan"`
    Prescription   string `json:"prescription"`
    Notes          string `json:"notes"`
}

type UpdateMedicalRecordRequest struct {
    VisitDate      string `json:"visit_date" binding:"omitempty,datetime=2006-01-02"`
    ChiefComplaint string `json:"chief_complaint"`
    Diagnosis      string `json:"diagnosis"`
    Treatment      string `json:"treatment"`
    TreatmentPlan  string `json:"treatment_plan"`
    Prescription   string `json:"prescription"`
    Notes          string `json:"notes"`
}

type MedicalRecordResponse struct {
    ID              uint   `json:"id"`
    PatientID       uint   `json:"patient_id"`
    DoctorID        uint   `json:"doctor_id"`
    VisitDate       string `json:"visit_date"`
    ChiefComplaint  string `json:"chief_complaint"`
    Diagnosis       string `json:"diagnosis"`
    Treatment       string `json:"treatment"`
    TreatmentPlan   string `json:"treatment_plan"`
    Prescription    string `json:"prescription"`
    Notes           string `json:"notes"`
    CreatedAt       string `json:"created_at"`
    UpdatedAt       string `json:"updated_at"`
}
