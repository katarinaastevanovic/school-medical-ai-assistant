package repository

import (
	"medical-service/database"
	"medical-service/models"
)

type MedicalRepository interface {
	GetOrCreateDoctor(userID uint) (models.Doctor, error)
	GetOrCreatePatient(userID uint) (models.Patient, error)
	GetAllPatients() ([]models.Patient, error)
	GetOrCreateMedicalRecord(patientID uint) (models.MedicalRecord, error)
	GetMedicalRecord(patientID uint) (models.MedicalRecord, error)
	CreateExaminationRequest(req *models.ExaminationRequest) error
	GetExaminationRequests(doctorID uint) ([]models.ExaminationRequest, error)
	GetMyExaminationRequests(patientID uint) ([]models.ExaminationRequest, error)
	UpdateRequestStatus(requestID uint, status models.TypeOfRequest) error
	CreateExamination(exam *models.Examination) error
	CreateCertificate(cert *models.MedicalCertificate) error
	GetMyCertificates(patientID uint) ([]models.MedicalCertificate, error)
	GetRequestByID(requestID uint) (models.ExaminationRequest, error)
	GetPatientByID(patientID uint) (models.Patient, error)
	GetPatientByUserID(userID uint) (models.Patient, error)
	GetCertificatesByPatientID(patientID uint) ([]models.MedicalCertificate, error)
	UpdateDoctorWorkingHours(userID uint, hours string) error
	GetAllDoctors() ([]models.Doctor, error)
}

type medicalRepository struct{}

func NewMedicalRepository() MedicalRepository {
	return &medicalRepository{}
}

func (r *medicalRepository) GetOrCreateDoctor(userID uint) (models.Doctor, error) {
	var doctor models.Doctor
	err := database.DB.Where("user_id = ?", userID).First(&doctor).Error
	if err == nil {
		return doctor, nil
	}
	doctor = models.Doctor{UserID: userID}
	err = database.DB.Create(&doctor).Error
	return doctor, err
}

func (r *medicalRepository) GetOrCreatePatient(userID uint) (models.Patient, error) {
	var patient models.Patient
	err := database.DB.Where("user_id = ?", userID).First(&patient).Error
	if err == nil {
		return patient, nil
	}

	patient = models.Patient{UserID: userID}
	if err := database.DB.Create(&patient).Error; err != nil {
		return models.Patient{}, err
	}
	return patient, nil
}

func (r *medicalRepository) GetAllPatients() ([]models.Patient, error) {
	var patients []models.Patient
	err := database.DB.Find(&patients).Error
	return patients, err
}

func (r *medicalRepository) GetOrCreateMedicalRecord(patientID uint) (models.MedicalRecord, error) {
	var record models.MedicalRecord
	err := database.DB.Where("patient_id = ?", patientID).First(&record).Error
	if err == nil {
		return record, nil
	}

	record = models.MedicalRecord{PatientID: patientID}
	if err := database.DB.Create(&record).Error; err != nil {
		return models.MedicalRecord{}, err
	}
	return record, nil
}

func (r *medicalRepository) GetMedicalRecord(patientID uint) (models.MedicalRecord, error) {
	var record models.MedicalRecord
	err := database.DB.
		Preload("Examinations").
		Where("patient_id = ?", patientID).
		First(&record).Error
	return record, err
}

func (r *medicalRepository) CreateExaminationRequest(req *models.ExaminationRequest) error {
	return database.DB.Create(req).Error
}

func (r *medicalRepository) GetExaminationRequests(doctorID uint) ([]models.ExaminationRequest, error) {
	var requests []models.ExaminationRequest
	err := database.DB.Where("doctor_id = ?", doctorID).Order("created_at desc").Find(&requests).Error
	return requests, err
}

func (r *medicalRepository) GetMyExaminationRequests(patientID uint) ([]models.ExaminationRequest, error) {
	var requests []models.ExaminationRequest
	err := database.DB.Where("patient_id = ?", patientID).Order("created_at desc").Find(&requests).Error
	return requests, err
}

func (r *medicalRepository) UpdateRequestStatus(requestID uint, status models.TypeOfRequest) error {
	return database.DB.Model(&models.ExaminationRequest{}).
		Where("id = ?", requestID).
		Update("status", status).Error
}

func (r *medicalRepository) CreateExamination(exam *models.Examination) error {
	return database.DB.Create(exam).Error
}

func (r *medicalRepository) CreateCertificate(cert *models.MedicalCertificate) error {
	return database.DB.Create(cert).Error
}

func (r *medicalRepository) GetMyCertificates(patientID uint) ([]models.MedicalCertificate, error) {
	var certs []models.MedicalCertificate
	err := database.DB.Where("patient_id = ?", patientID).Order("created_at desc").Find(&certs).Error
	return certs, err
}

func (r *medicalRepository) GetRequestByID(requestID uint) (models.ExaminationRequest, error) {
	var req models.ExaminationRequest
	err := database.DB.First(&req, requestID).Error
	return req, err
}

func (r *medicalRepository) GetPatientByID(patientID uint) (models.Patient, error) {
	var patient models.Patient
	err := database.DB.First(&patient, patientID).Error
	return patient, err
}

func (r *medicalRepository) GetCertificatesByPatientID(patientID uint) ([]models.MedicalCertificate, error) {
	var certs []models.MedicalCertificate
	err := database.DB.Where("patient_id = ?", patientID).Order("created_at desc").Find(&certs).Error
	return certs, err
}

func (r *medicalRepository) GetPatientByUserID(userID uint) (models.Patient, error) {
	var patient models.Patient
	err := database.DB.Where("user_id = ?", userID).First(&patient).Error
	return patient, err
}

func (r *medicalRepository) UpdateDoctorWorkingHours(userID uint, hours string) error {
	return database.DB.Model(&models.Doctor{}).
		Where("user_id = ?", userID).
		Update("working_hours", hours).Error
}

func (r *medicalRepository) GetAllDoctors() ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := database.DB.Find(&doctors).Error
	return doctors, err
}
