package handlers

import (
	"encoding/json"
	"medical-service/models"
	"medical-service/services"
	"net/http"
	"strconv"
	"time"
)

type MedicalHandler struct {
	Service services.MedicalService
}

func (h *MedicalHandler) CreateExaminationRequest(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var body struct {
		DoctorID uint                     `json:"doctorID"`
		Type     models.TypeOfExamination `json:"type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateExaminationRequest(uint(userID), body.DoctorID, body.Type); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MedicalHandler) GetExaminationRequests(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	requests, err := h.Service.GetExaminationRequests(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// enrich sa UserID
	type EnrichedRequest struct {
		ID              uint   `json:"id"`
		PatientID       uint   `json:"patientID"`
		PatientUserID   uint   `json:"patientUserID"`
		DoctorID        uint   `json:"doctorID"`
		Type            string `json:"type"`
		Status          string `json:"status"`
		MedicalRecordID uint   `json:"medicalRecordID"`
	}

	result := make([]EnrichedRequest, 0)
	for _, req := range requests {
		patient, err := h.Service.GetPatientByID(req.PatientID)
		if err != nil {
			continue
		}
		result = append(result, EnrichedRequest{
			ID:              req.ID,
			PatientID:       req.PatientID,
			PatientUserID:   patient.UserID,
			DoctorID:        req.DoctorID,
			Type:            string(req.Type),
			Status:          string(req.Status),
			MedicalRecordID: req.MedicalRecordID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *MedicalHandler) GetMyExaminationRequests(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	requests, err := h.Service.GetMyExaminationRequests(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}

func (h *MedicalHandler) UpdateRequestStatus(w http.ResponseWriter, r *http.Request) {
	requestIDStr := r.PathValue("id")
	requestID, err := strconv.ParseUint(requestIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid request ID", http.StatusBadRequest)
		return
	}

	var body struct {
		Status models.TypeOfRequest `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateRequestStatus(uint(requestID), body.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MedicalHandler) CreateExamination(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	requestIDStr := r.PathValue("id")
	requestID, err := strconv.ParseUint(requestIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid request ID", http.StatusBadRequest)
		return
	}

	var body struct {
		Diagnosis string `json:"diagnosis"`
		Therapy   string `json:"therapy"`
		Note      string `json:"note"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateExamination(uint(userID), uint(requestID), body.Diagnosis, body.Therapy, body.Note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MedicalHandler) CreateCertificate(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var body struct {
		PatientUserID uint                     `json:"patientUserID"`
		Type          models.TypeOfCertificate `json:"type"`
		Note          string                   `json:"note"`
		Date          time.Time                `json:"date"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateCertificate(uint(userID), body.PatientUserID, body.Type, body.Note, body.Date); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MedicalHandler) GetMyCertificates(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	certs, err := h.Service.GetMyCertificates(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(certs)
}

func (h *MedicalHandler) GetMyMedicalRecord(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	record, err := h.Service.GetMyMedicalRecord(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func (h *MedicalHandler) GetAllPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := h.Service.GetAllPatients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}

func (h *MedicalHandler) GetCertificatesByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.PathValue("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	certs, err := h.Service.GetCertificatesByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(certs)
}

func (h *MedicalHandler) GetMedicalRecordByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.PathValue("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	record, err := h.Service.GetMedicalRecordByUserID(uint(userID))
	if err != nil {
		switch err.Error() {
		case "patient not found", "medical record not found":
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func (h *MedicalHandler) UpdateMyWorkingHours(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var body struct {
		WorkingHours string `json:"workingHours"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateDoctorWorkingHours(uint(userID), body.WorkingHours); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MedicalHandler) GetAllDoctorsEnriched(w http.ResponseWriter, r *http.Request) {
	doctors, err := h.Service.GetAllDoctorsEnriched()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doctors)
}

func (h *MedicalHandler) GetChatbotContext(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetChatbotContext()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
