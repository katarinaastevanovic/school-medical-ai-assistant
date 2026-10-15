package services

import (
	"encoding/json"
	"errors"
	"medical-service/models"
	"medical-service/repository"
	"net/http"
	"time"
)

type MedicalService struct {
	Repo repository.MedicalRepository
}

func (s *MedicalService) CreateExaminationRequest(userID uint, doctorUserID uint, examType models.TypeOfExamination) error {
	patient, err := s.Repo.GetOrCreatePatient(userID)
	if err != nil {
		return err
	}

	record, err := s.Repo.GetOrCreateMedicalRecord(patient.ID)
	if err != nil {
		return err
	}

	doctor, err := s.Repo.GetOrCreateDoctor(doctorUserID)
	if err != nil {
		return err
	}

	req := &models.ExaminationRequest{
		PatientID:       patient.ID,
		DoctorID:        doctor.ID,
		Type:            examType,
		Status:          models.REQUESTED,
		MedicalRecordID: record.ID,
	}

	return s.Repo.CreateExaminationRequest(req)
}

func (s *MedicalService) GetExaminationRequests(userID uint) ([]models.ExaminationRequest, error) {
	doctor, err := s.Repo.GetOrCreateDoctor(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetExaminationRequests(doctor.ID)
}

func (s *MedicalService) GetMyExaminationRequests(userID uint) ([]models.ExaminationRequest, error) {
	patient, err := s.Repo.GetOrCreatePatient(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetMyExaminationRequests(patient.ID)
}

func (s *MedicalService) UpdateRequestStatus(requestID uint, status models.TypeOfRequest) error {
	return s.Repo.UpdateRequestStatus(requestID, status)
}

func (s *MedicalService) CreateExamination(doctorUserID, requestID uint, diagnosis, therapy, note string) error {
	req, err := s.Repo.GetRequestByID(requestID)
	if err != nil {
		return errors.New("request not found")
	}

	if req.Status != models.APPROVED {
		return errors.New("request must be approved before examination")
	}

	exam := &models.Examination{
		RequestID:       requestID,
		MedicalRecordID: req.MedicalRecordID,
		Diagnosis:       diagnosis,
		Therapy:         therapy,
		Note:            note,
	}

	if err := s.Repo.CreateExamination(exam); err != nil {
		return err
	}

	return s.Repo.UpdateRequestStatus(requestID, models.APPROVED)
}

func (s *MedicalService) CreateCertificate(doctorUserID, patientUserID uint, certType models.TypeOfCertificate, note string, date time.Time) error {
	doctor, err := s.Repo.GetOrCreateDoctor(doctorUserID)
	if err != nil {
		return err
	}

	patient, err := s.Repo.GetOrCreatePatient(patientUserID)
	if err != nil {
		return err
	}

	cert := &models.MedicalCertificate{
		PatientID: patient.ID,
		DoctorID:  doctor.ID,
		Date:      date,
		Type:      certType,
		Note:      note,
	}

	return s.Repo.CreateCertificate(cert)
}

func (s *MedicalService) GetMyCertificates(userID uint) ([]models.MedicalCertificate, error) {
	patient, err := s.Repo.GetOrCreatePatient(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetMyCertificates(patient.ID)
}

func (s *MedicalService) GetMyMedicalRecord(userID uint) (models.MedicalRecord, error) {
	patient, err := s.Repo.GetOrCreatePatient(userID)
	if err != nil {
		return models.MedicalRecord{}, err
	}

	record, err := s.Repo.GetMedicalRecord(patient.ID)
	if err != nil {
		return models.MedicalRecord{}, err
	}

	return record, nil
}

func (s *MedicalService) GetAllPatients() ([]models.Patient, error) {
	return s.Repo.GetAllPatients()
}

func (s *MedicalService) GetPatientByID(patientID uint) (models.Patient, error) {
	return s.Repo.GetPatientByID(patientID)
}

func (s *MedicalService) GetCertificatesByUserID(userID uint) ([]models.MedicalCertificate, error) {
	patient, err := s.Repo.GetOrCreatePatient(userID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetCertificatesByPatientID(patient.ID)
}

func (s *MedicalService) GetMedicalRecordByPatientID(patientID uint) (models.MedicalRecord, error) {
	return s.Repo.GetMedicalRecord(patientID)
}

func (s *MedicalService) GetMedicalRecordByUserID(userID uint) (models.MedicalRecord, error) {
	patient, err := s.Repo.GetPatientByUserID(userID)
	if err != nil {
		return models.MedicalRecord{}, errors.New("patient not found")
	}

	record, err := s.Repo.GetMedicalRecord(patient.ID)
	if err != nil {
		return models.MedicalRecord{}, errors.New("medical record not found")
	}

	return record, nil
}

func (s *MedicalService) UpdateDoctorWorkingHours(userID uint, hours string) error {
	return s.Repo.UpdateDoctorWorkingHours(userID, hours)
}

func (s *MedicalService) GetAllDoctorsEnriched() ([]map[string]interface{}, error) {
	doctors, err := s.Repo.GetAllDoctors()
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("http://auth-service:8081/members/doctors")
	if err != nil {
		return nil, errors.New("failed to reach auth service")
	}
	defer resp.Body.Close()

	var allDoctors []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&allDoctors); err != nil {
		return nil, errors.New("failed to parse auth service response")
	}

	hoursMap := make(map[uint]string)
	for _, d := range doctors {
		hoursMap[d.UserID] = d.WorkingHours
	}

	result := make([]map[string]interface{}, 0)
	for _, d := range allDoctors {
		idFloat, ok := d["id"].(float64)
		if !ok {
			continue
		}
		userID := uint(idFloat)
		result = append(result, map[string]interface{}{
			"userID":       userID,
			"firstName":    d["firstName"],
			"lastName":     d["lastName"],
			"workingHours": hoursMap[userID],
		})
	}

	return result, nil
}

func (s *MedicalService) GetChatbotContext() (map[string]interface{}, error) {
	doctors, err := s.GetAllDoctorsEnriched()
	if err != nil {
		return nil, err
	}

	doctorInfo := make([]map[string]interface{}, 0, len(doctors))
	for _, d := range doctors {
		doctorInfo = append(doctorInfo, map[string]interface{}{
			"firstName":    d["firstName"],
			"lastName":     d["lastName"],
			"workingHours": d["workingHours"],
		})
	}

	return map[string]interface{}{
		"doctors":          doctorInfo,
		"examinationTypes": []string{"GENERAL", "SPECIALIST"},
		"certificateTypes": []string{"REGULAR", "PE", "EXCURSION", "SICKNESS"},
	}, nil
}
