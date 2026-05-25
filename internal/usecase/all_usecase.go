package usecase

import (
    "context"
    "ehr-api/internal/dto"
    "ehr-api/internal/entity"
    "ehr-api/internal/repository"
    "time"
)

// ---------- Allergy ----------
type AllergyUsecase interface {
    Create(ctx context.Context, req dto.CreateAllergyRequest) (dto.AllergyResponse, error)
    FindAll(ctx context.Context) ([]dto.AllergyResponse, error)
    FindByID(ctx context.Context, id uint) (dto.AllergyResponse, error)
    Update(ctx context.Context, id uint, req dto.UpdateAllergyRequest) (dto.AllergyResponse, error)
    Delete(ctx context.Context, id uint) error
}

type allergyUsecase struct {
    repo repository.AllergyRepository
}

func NewAllergyUsecase(r repository.AllergyRepository) AllergyUsecase { return &allergyUsecase{repo: r} }

func (u *allergyUsecase) Create(ctx context.Context, req dto.CreateAllergyRequest) (dto.AllergyResponse, error) {
    noted, _ := time.Parse("2006-01-02", req.NotedDate)
    model := entity.Allergy{
        PatientID: req.PatientID,
        Allergen:  req.Allergen,
        Severity:  req.Severity,
        Reaction:  req.Reaction,
        NotedDate: noted,
    }
    if err := u.repo.Create(ctx, &model); err != nil { return dto.AllergyResponse{}, err }
    return mapAllergyToResponse(model), nil
}

func (u *allergyUsecase) FindAll(ctx context.Context) ([]dto.AllergyResponse, error) {
    list, err := u.repo.FindAll(ctx)
    if err != nil { return nil, err }
    res := make([]dto.AllergyResponse, len(list))
    for i, a := range list { res[i] = mapAllergyToResponse(a) }
    return res, nil
}

func (u *allergyUsecase) FindByID(ctx context.Context, id uint) (dto.AllergyResponse, error) {
    a, err := u.repo.FindByID(ctx, id)
    if err != nil { return dto.AllergyResponse{}, err }
    return mapAllergyToResponse(*a), nil
}

func (u *allergyUsecase) Update(ctx context.Context, id uint, req dto.UpdateAllergyRequest) (dto.AllergyResponse, error) {
    a, err := u.repo.FindByID(ctx, id)
    if err != nil { return dto.AllergyResponse{}, err }
    if req.Allergen != "" { a.Allergen = req.Allergen }
    if req.Severity != "" { a.Severity = req.Severity }
    if req.Reaction != "" { a.Reaction = req.Reaction }
    if req.NotedDate != "" { nd, _ := time.Parse("2006-01-02", req.NotedDate); a.NotedDate = nd }
    if err := u.repo.Update(ctx, a); err != nil { return dto.AllergyResponse{}, err }
    return mapAllergyToResponse(*a), nil
}

func (u *allergyUsecase) Delete(ctx context.Context, id uint) error { return u.repo.Delete(ctx, id) }

func mapAllergyToResponse(a entity.Allergy) dto.AllergyResponse {
    return dto.AllergyResponse{ID: a.ID, PatientID: a.PatientID, Allergen: a.Allergen, Severity: a.Severity, Reaction: a.Reaction, NotedDate: a.NotedDate.Format("2006-01-02"), CreatedAt: a.CreatedAt.Format(time.RFC3339), UpdatedAt: a.UpdatedAt.Format(time.RFC3339)}
}

// ---------- Appointment ----------
type AppointmentUsecase interface {
    Create(ctx context.Context, req dto.CreateAppointmentRequest) (dto.AppointmentResponse, error)
    FindAll(ctx context.Context) ([]dto.AppointmentResponse, error)
    FindByID(ctx context.Context, id uint) (dto.AppointmentResponse, error)
    Update(ctx context.Context, id uint, req dto.UpdateAppointmentRequest) (dto.AppointmentResponse, error)
    Delete(ctx context.Context, id uint) error
}

type appointmentUsecase struct { repo repository.AppointmentRepository }
func NewAppointmentUsecase(r repository.AppointmentRepository) AppointmentUsecase { return &appointmentUsecase{repo: r} }
func (u *appointmentUsecase) Create(ctx context.Context, req dto.CreateAppointmentRequest) (dto.AppointmentResponse, error) {
    dt, _ := time.Parse(time.RFC3339, req.AppointmentDate)
    m := entity.Appointment{PatientID: req.PatientID, DoctorID: req.DoctorID, AppointmentDate: dt, Status: req.Status, Reason: req.Reason, Notes: req.Notes}
    if err := u.repo.Create(ctx, &m); err != nil { return dto.AppointmentResponse{}, err }
    return mapAppointment(m), nil
}
func (u *appointmentUsecase) FindAll(ctx context.Context) ([]dto.AppointmentResponse, error) {
    list, err := u.repo.FindAll(ctx); if err != nil { return nil, err }
    res := make([]dto.AppointmentResponse, len(list))
    for i, a := range list { res[i] = mapAppointment(a) }
    return res, nil
}
func (u *appointmentUsecase) FindByID(ctx context.Context, id uint) (dto.AppointmentResponse, error) {
    a, err := u.repo.FindByID(ctx, id); if err != nil { return dto.AppointmentResponse{}, err }
    return mapAppointment(*a), nil
}
func (u *appointmentUsecase) Update(ctx context.Context, id uint, req dto.UpdateAppointmentRequest) (dto.AppointmentResponse, error) {
    a, err := u.repo.FindByID(ctx, id); if err != nil { return dto.AppointmentResponse{}, err }
    if req.AppointmentDate != "" { dt, _ := time.Parse(time.RFC3339, req.AppointmentDate); a.AppointmentDate = dt }
    if req.Status != "" { a.Status = req.Status }
    if req.Reason != "" { a.Reason = req.Reason }
    if req.Notes != "" { a.Notes = req.Notes }
    if err := u.repo.Update(ctx, a); err != nil { return dto.AppointmentResponse{}, err }
    return mapAppointment(*a), nil
}
func (u *appointmentUsecase) Delete(ctx context.Context, id uint) error { return u.repo.Delete(ctx, id) }
func mapAppointment(a entity.Appointment) dto.AppointmentResponse {
    return dto.AppointmentResponse{ID: a.ID, PatientID: a.PatientID, DoctorID: a.DoctorID, AppointmentDate: a.AppointmentDate.Format(time.RFC3339), Status: a.Status, Reason: a.Reason, Notes: a.Notes, CreatedAt: a.CreatedAt.Format(time.RFC3339), UpdatedAt: a.UpdatedAt.Format(time.RFC3339)}
}

// ---------- Medication ----------
type MedicationUsecase interface {
    Create(ctx context.Context, req dto.CreateMedicationRequest) (dto.MedicationResponse, error)
    FindAll(ctx context.Context) ([]dto.MedicationResponse, error)
    FindByID(ctx context.Context, id uint) (dto.MedicationResponse, error)
    Update(ctx context.Context, id uint, req dto.UpdateMedicationRequest) (dto.MedicationResponse, error)
    Delete(ctx context.Context, id uint) error
}

type medicationUsecase struct { repo repository.MedicationRepository }
func NewMedicationUsecase(r repository.MedicationRepository) MedicationUsecase { return &medicationUsecase{repo: r} }
func (u *medicationUsecase) Create(ctx context.Context, req dto.CreateMedicationRequest) (dto.MedicationResponse, error) {
    m := entity.Medication{MedicalRecordID: req.MedicalRecordID, Name: req.Name, Dosage: req.Dosage, Frequency: req.Frequency, Duration: req.Duration, Notes: req.Notes}
    if err := u.repo.Create(ctx, &m); err != nil { return dto.MedicationResponse{}, err }
    return mapMedication(m), nil
}
func (u *medicationUsecase) FindAll(ctx context.Context) ([]dto.MedicationResponse, error) {
    list, err := u.repo.FindAll(ctx); if err != nil { return nil, err }
    out := make([]dto.MedicationResponse, len(list))
    for i, m := range list { out[i] = mapMedication(m) }
    return out, nil
}
func (u *medicationUsecase) FindByID(ctx context.Context, id uint) (dto.MedicationResponse, error) {
    m, err := u.repo.FindByID(ctx, id); if err != nil { return dto.MedicationResponse{}, err }
    return mapMedication(*m), nil
}
func (u *medicationUsecase) Update(ctx context.Context, id uint, req dto.UpdateMedicationRequest) (dto.MedicationResponse, error) {
    m, err := u.repo.FindByID(ctx, id); if err != nil { return dto.MedicationResponse{}, err }
    if req.Name != "" { m.Name = req.Name }
    if req.Dosage != "" { m.Dosage = req.Dosage }
    if req.Frequency != "" { m.Frequency = req.Frequency }
    if req.Duration != "" { m.Duration = req.Duration }
    if req.Notes != "" { m.Notes = req.Notes }
    if err := u.repo.Update(ctx, m); err != nil { return dto.MedicationResponse{}, err }
    return mapMedication(*m), nil
}
func (u *medicationUsecase) Delete(ctx context.Context, id uint) error { return u.repo.Delete(ctx, id) }
func mapMedication(m entity.Medication) dto.MedicationResponse {
    return dto.MedicationResponse{ID: m.ID, MedicalRecordID: m.MedicalRecordID, Name: m.Name, Dosage: m.Dosage, Frequency: m.Frequency, Duration: m.Duration, Notes: m.Notes, CreatedAt: m.CreatedAt.Format(time.RFC3339), UpdatedAt: m.UpdatedAt.Format(time.RFC3339)}
}

// ---------- LabResult ----------
type LabResultUsecase interface {
    Create(ctx context.Context, req dto.CreateLabResultRequest) (dto.LabResultResponse, error)
    FindAll(ctx context.Context) ([]dto.LabResultResponse, error)
    FindByID(ctx context.Context, id uint) (dto.LabResultResponse, error)
    Update(ctx context.Context, id uint, req dto.UpdateLabResultRequest) (dto.LabResultResponse, error)
    Delete(ctx context.Context, id uint) error
}

type labResultUsecase struct { repo repository.LabResultRepository }
func NewLabResultUsecase(r repository.LabResultRepository) LabResultUsecase { return &labResultUsecase{repo: r} }
func (u *labResultUsecase) Create(ctx context.Context, req dto.CreateLabResultRequest) (dto.LabResultResponse, error) {
    td, _ := time.Parse(time.RFC3339, req.TestDate)
    m := entity.LabResult{MedicalRecordID: req.MedicalRecordID, TestName: req.TestName, Result: req.Result, Unit: req.Unit, ReferenceRange: req.ReferenceRange, TestDate: td, Remarks: req.Remarks}
    if err := u.repo.Create(ctx, &m); err != nil { return dto.LabResultResponse{}, err }
    return mapLabResult(m), nil
}
func (u *labResultUsecase) FindAll(ctx context.Context) ([]dto.LabResultResponse, error) {
    list, err := u.repo.FindAll(ctx); if err != nil { return nil, err }
    out := make([]dto.LabResultResponse, len(list))
    for i, lr := range list { out[i] = mapLabResult(lr) }
    return out, nil
}
func (u *labResultUsecase) FindByID(ctx context.Context, id uint) (dto.LabResultResponse, error) {
    lr, err := u.repo.FindByID(ctx, id); if err != nil { return dto.LabResultResponse{}, err }
    return mapLabResult(*lr), nil
}
func (u *labResultUsecase) Update(ctx context.Context, id uint, req dto.UpdateLabResultRequest) (dto.LabResultResponse, error) {
    lr, err := u.repo.FindByID(ctx, id); if err != nil { return dto.LabResultResponse{}, err }
    if req.TestName != "" { lr.TestName = req.TestName }
    if req.Result != "" { lr.Result = req.Result }
    if req.Unit != "" { lr.Unit = req.Unit }
    if req.ReferenceRange != "" { lr.ReferenceRange = req.ReferenceRange }
    if req.TestDate != "" { td, _ := time.Parse(time.RFC3339, req.TestDate); lr.TestDate = td }
    if req.Remarks != "" { lr.Remarks = req.Remarks }
    if err := u.repo.Update(ctx, lr); err != nil { return dto.LabResultResponse{}, err }
    return mapLabResult(*lr), nil
}
func (u *labResultUsecase) Delete(ctx context.Context, id uint) error { return u.repo.Delete(ctx, id) }
func mapLabResult(lr entity.LabResult) dto.LabResultResponse {
    return dto.LabResultResponse{ID: lr.ID, MedicalRecordID: lr.MedicalRecordID, TestName: lr.TestName, Result: lr.Result, Unit: lr.Unit, ReferenceRange: lr.ReferenceRange, TestDate: lr.TestDate.Format(time.RFC3339), Remarks: lr.Remarks, CreatedAt: lr.CreatedAt.Format(time.RFC3339), UpdatedAt: lr.UpdatedAt.Format(time.RFC3339)}
}

// ---------- VitalSign ----------
type VitalSignUsecase interface {
    Create(ctx context.Context, req dto.CreateVitalSignRequest) (dto.VitalSignResponse, error)
    FindAll(ctx context.Context) ([]dto.VitalSignResponse, error)
    FindByID(ctx context.Context, id uint) (dto.VitalSignResponse, error)
    Update(ctx context.Context, id uint, req dto.UpdateVitalSignRequest) (dto.VitalSignResponse, error)
    Delete(ctx context.Context, id uint) error
}

type vitalSignUsecase struct { repo repository.VitalSignRepository }
func NewVitalSignUsecase(r repository.VitalSignRepository) VitalSignUsecase { return &vitalSignUsecase{repo: r} }
func (u *vitalSignUsecase) Create(ctx context.Context, req dto.CreateVitalSignRequest) (dto.VitalSignResponse, error) {
    ra, _ := time.Parse(time.RFC3339, req.RecordedAt)
    m := entity.VitalSign{MedicalRecordID: req.MedicalRecordID, BloodPressure: req.BloodPressure, HeartRate: req.HeartRate, Temperature: req.Temperature, RespiratoryRate: req.RespiratoryRate, OxygenSaturation: req.OxygenSaturation, Weight: req.Weight, Height: req.Height, RecordedAt: ra}
    if err := u.repo.Create(ctx, &m); err != nil { return dto.VitalSignResponse{}, err }
    return mapVitalSign(m), nil
}
func (u *vitalSignUsecase) FindAll(ctx context.Context) ([]dto.VitalSignResponse, error) {
    list, err := u.repo.FindAll(ctx); if err != nil { return nil, err }
    out := make([]dto.VitalSignResponse, len(list))
    for i, v := range list { out[i] = mapVitalSign(v) }
    return out, nil
}
func (u *vitalSignUsecase) FindByID(ctx context.Context, id uint) (dto.VitalSignResponse, error) {
    v, err := u.repo.FindByID(ctx, id); if err != nil { return dto.VitalSignResponse{}, err }
    return mapVitalSign(*v), nil
}
func (u *vitalSignUsecase) Update(ctx context.Context, id uint, req dto.UpdateVitalSignRequest) (dto.VitalSignResponse, error) {
    v, err := u.repo.FindByID(ctx, id); if err != nil { return dto.VitalSignResponse{}, err }
    if req.BloodPressure != "" { v.BloodPressure = req.BloodPressure }
    if req.HeartRate != nil { v.HeartRate = *req.HeartRate }
    if req.Temperature != nil { v.Temperature = *req.Temperature }
    if req.RespiratoryRate != nil { v.RespiratoryRate = *req.RespiratoryRate }
    if req.OxygenSaturation != nil { v.OxygenSaturation = *req.OxygenSaturation }
    if req.Weight != nil { v.Weight = *req.Weight }
    if req.Height != nil { v.Height = *req.Height }
    if req.RecordedAt != "" { ra, _ := time.Parse(time.RFC3339, req.RecordedAt); v.RecordedAt = ra }
    if err := u.repo.Update(ctx, v); err != nil { return dto.VitalSignResponse{}, err }
    return mapVitalSign(*v), nil
}
func (u *vitalSignUsecase) Delete(ctx context.Context, id uint) error { return u.repo.Delete(ctx, id) }
func mapVitalSign(v entity.VitalSign) dto.VitalSignResponse {
    return dto.VitalSignResponse{ID: v.ID, MedicalRecordID: v.MedicalRecordID, BloodPressure: v.BloodPressure, HeartRate: v.HeartRate, Temperature: v.Temperature, RespiratoryRate: v.RespiratoryRate, OxygenSaturation: v.OxygenSaturation, Weight: v.Weight, Height: v.Height, RecordedAt: v.RecordedAt.Format(time.RFC3339), CreatedAt: v.CreatedAt.Format(time.RFC3339), UpdatedAt: v.UpdatedAt.Format(time.RFC3339)}
}

// ---------- User ----------
type UserUsecase interface {
    Create(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error)
    FindAll(ctx context.Context) ([]dto.UserResponse, error)
    FindByID(ctx context.Context, id uint) (dto.UserResponse, error)
    Update(ctx context.Context, id uint, req dto.UpdateUserRequest) (dto.UserResponse, error)
    Delete(ctx context.Context, id uint) error
}

type userUsecase struct { repo repository.UserRepository }
func NewUserUsecase(r repository.UserRepository) UserUsecase { return &userUsecase{repo: r} }
func (u *userUsecase) Create(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error) {
    m := entity.User{Name: req.Name, Username: req.Username, Password: req.Password, Role: req.Role, Specialization: req.Specialization, Phone: req.Phone, Email: req.Email}
    if err := u.repo.Create(ctx, &m); err != nil { return dto.UserResponse{}, err }
    return mapUser(m), nil
}
func (u *userUsecase) FindAll(ctx context.Context) ([]dto.UserResponse, error) {
    list, err := u.repo.FindAll(ctx); if err != nil { return nil, err }
    out := make([]dto.UserResponse, len(list))
    for i, usr := range list { out[i] = mapUser(usr) }
    return out, nil
}
func (u *userUsecase) FindByID(ctx context.Context, id uint) (dto.UserResponse, error) {
    usr, err := u.repo.FindByID(ctx, id); if err != nil { return dto.UserResponse{}, err }
    return mapUser(*usr), nil
}
func (u *userUsecase) Update(ctx context.Context, id uint, req dto.UpdateUserRequest) (dto.UserResponse, error) {
    usr, err := u.repo.FindByID(ctx, id); if err != nil { return dto.UserResponse{}, err }
    if req.Name != nil { usr.Name = *req.Name }
    if req.Username != nil { usr.Username = *req.Username }
    if req.Password != nil { usr.Password = *req.Password }
    if req.Role != nil { usr.Role = *req.Role }
    if req.Specialization != nil { usr.Specialization = *req.Specialization }
    if req.Phone != nil { usr.Phone = *req.Phone }
    if req.Email != nil { usr.Email = *req.Email }
    if err := u.repo.Update(ctx, usr); err != nil { return dto.UserResponse{}, err }
    return mapUser(*usr), nil
}
func (u *userUsecase) Delete(ctx context.Context, id uint) error { return u.repo.Delete(ctx, id) }
func mapUser(usr entity.User) dto.UserResponse {
    return dto.UserResponse{ID: usr.ID, Name: usr.Name, Username: usr.Username, Role: usr.Role, Specialization: usr.Specialization, Phone: usr.Phone, Email: usr.Email, CreatedAt: usr.CreatedAt.Format(time.RFC3339), UpdatedAt: usr.UpdatedAt.Format(time.RFC3339)}
}

// ---------- MedicalRecord ----------
type MedicalRecordUsecase interface {
    Create(ctx context.Context, req dto.CreateMedicalRecordRequest) (dto.MedicalRecordResponse, error)
    FindAll(ctx context.Context) ([]dto.MedicalRecordResponse, error)
    FindByID(ctx context.Context, id uint) (dto.MedicalRecordResponse, error)
    Update(ctx context.Context, id uint, req dto.UpdateMedicalRecordRequest) (dto.MedicalRecordResponse, error)
    Delete(ctx context.Context, id uint) error
}

type medicalRecordUsecase struct { repo repository.MedicalRecordRepository }
func NewMedicalRecordUsecase(r repository.MedicalRecordRepository) MedicalRecordUsecase { return &medicalRecordUsecase{repo: r} }
func (u *medicalRecordUsecase) Create(ctx context.Context, req dto.CreateMedicalRecordRequest) (dto.MedicalRecordResponse, error) {
    vd, _ := time.Parse("2006-01-02", req.VisitDate)
    m := entity.MedicalRecord{PatientID: req.PatientID, DoctorID: req.DoctorID, VisitDate: vd, ChiefComplaint: req.ChiefComplaint, Diagnosis: req.Diagnosis, Treatment: req.Treatment, TreatmentPlan: req.TreatmentPlan, Prescription: req.Prescription, Notes: req.Notes}
    if err := u.repo.Create(ctx, &m); err != nil { return dto.MedicalRecordResponse{}, err }
    return mapMedicalRecord(m), nil
}
func (u *medicalRecordUsecase) FindAll(ctx context.Context) ([]dto.MedicalRecordResponse, error) {
    list, err := u.repo.FindAll(ctx); if err != nil { return nil, err }
    out := make([]dto.MedicalRecordResponse, len(list))
    for i, mr := range list { out[i] = mapMedicalRecord(mr) }
    return out, nil
}
func (u *medicalRecordUsecase) FindByID(ctx context.Context, id uint) (dto.MedicalRecordResponse, error) {
    mr, err := u.repo.FindByID(ctx, id); if err != nil { return dto.MedicalRecordResponse{}, err }
    return mapMedicalRecord(*mr), nil
}
func (u *medicalRecordUsecase) Update(ctx context.Context, id uint, req dto.UpdateMedicalRecordRequest) (dto.MedicalRecordResponse, error) {
    mr, err := u.repo.FindByID(ctx, id); if err != nil { return dto.MedicalRecordResponse{}, err }
    if req.VisitDate != "" { vd, _ := time.Parse("2006-01-02", req.VisitDate); mr.VisitDate = vd }
    if req.ChiefComplaint != "" { mr.ChiefComplaint = req.ChiefComplaint }
    if req.Diagnosis != "" { mr.Diagnosis = req.Diagnosis }
    if req.Treatment != "" { mr.Treatment = req.Treatment }
    if req.TreatmentPlan != "" { mr.TreatmentPlan = req.TreatmentPlan }
    if req.Prescription != "" { mr.Prescription = req.Prescription }
    if req.Notes != "" { mr.Notes = req.Notes }
    if err := u.repo.Update(ctx, mr); err != nil { return dto.MedicalRecordResponse{}, err }
    return mapMedicalRecord(*mr), nil
}
func (u *medicalRecordUsecase) Delete(ctx context.Context, id uint) error { return u.repo.Delete(ctx, id) }
func mapMedicalRecord(mr entity.MedicalRecord) dto.MedicalRecordResponse {
    return dto.MedicalRecordResponse{ID: mr.ID, PatientID: mr.PatientID, DoctorID: mr.DoctorID, VisitDate: mr.VisitDate.Format("2006-01-02"), ChiefComplaint: mr.ChiefComplaint, Diagnosis: mr.Diagnosis, Treatment: mr.Treatment, TreatmentPlan: mr.TreatmentPlan, Prescription: mr.Prescription, Notes: mr.Notes, CreatedAt: mr.CreatedAt.Format(time.RFC3339), UpdatedAt: mr.UpdatedAt.Format(time.RFC3339)}
}
