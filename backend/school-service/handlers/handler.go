package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"school-service/dto"
	"school-service/models"
	"school-service/services"
)

type SchoolHandler struct {
	Service services.SchoolService
}

func (h *SchoolHandler) CreateClass(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var class models.SchoolClass

	err := json.NewDecoder(r.Body).Decode(&class)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateClass(&class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(class)
}

func (h *SchoolHandler) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	classes, err := h.Service.GetAllClasses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

func (h *SchoolHandler) GetStudentsFromClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	classIDStr := r.PathValue("id")
	classID, err := strconv.ParseUint(classIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	students, err := h.Service.GetStudentsFromClass(uint(classID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func (h *SchoolHandler) AddStudentToClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = h.Service.AddStudentToClass(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (h *SchoolHandler) RemoveStudentFromClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	studentIDStr := r.PathValue("studentId")
	studentID, err := strconv.ParseUint(studentIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	err = h.Service.RemoveStudentFromClass(uint(studentID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SchoolHandler) GetFreeStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	students, err := h.Service.GetFreeStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func (h *SchoolHandler) CreateSubject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var subject models.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateSubject(&subject); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subject)
}

func (h *SchoolHandler) GetAllSubjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	subjects, err := h.Service.GetAllSubjects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subjects)
}

func (h *SchoolHandler) UpdateSubject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateSubject(uint(id), body.Name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SchoolHandler) DeleteSubject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteSubject(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SchoolHandler) AssignTeacher(w http.ResponseWriter, r *http.Request) {
	var body struct {
		UserID    uint `json:"userID"`
		ClassID   uint `json:"classID"`
		SubjectID uint `json:"subjectID"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.AssignTeacher(body.UserID, body.ClassID, body.SubjectID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SchoolHandler) AddTeacherToSubject(w http.ResponseWriter, r *http.Request) {
	subjectIDStr := r.PathValue("id")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	var body struct {
		UserID uint `json:"userID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.AddTeacherToSubject(body.UserID, uint(subjectID)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SchoolHandler) GetTeachersBySubjectEnriched(w http.ResponseWriter, r *http.Request) {
	subjectIDStr := r.PathValue("id")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	teachers, err := h.Service.GetTeachersBySubjectEnriched(uint(subjectID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}

func (h *SchoolHandler) GetAllTeachersEnriched(w http.ResponseWriter, r *http.Request) {
	teachers, err := h.Service.GetAllTeachersEnriched()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}

func (h *SchoolHandler) GetSubjectsByClass(w http.ResponseWriter, r *http.Request) {
	classIDStr := r.PathValue("id")
	classID, err := strconv.ParseUint(classIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	subjects, err := h.Service.GetSubjectsByClass(uint(classID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subjects)
}

func (h *SchoolHandler) GetMySubjects(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	subjects, err := h.Service.GetSubjectsForTeacher(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subjects)
}

func (h *SchoolHandler) GetGradesAndAbsences(w http.ResponseWriter, r *http.Request) {
	studentUserIDStr := r.PathValue("userId")
	studentUserID, err := strconv.ParseUint(studentUserIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	subjectIDStr := r.PathValue("subjectId")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	data, err := h.Service.GetGradesAndAbsences(uint(studentUserID), uint(subjectID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *SchoolHandler) CreateGrade(w http.ResponseWriter, r *http.Request) {
	studentUserIDStr := r.PathValue("userId")
	studentUserID, err := strconv.ParseUint(studentUserIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	subjectIDStr := r.PathValue("subjectId")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	teacherUserIDStr := r.Header.Get("X-User-ID")
	teacherUserID, err := strconv.ParseUint(teacherUserIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	var body struct {
		Value int       `json:"value"`
		Date  time.Time `json:"date"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateGrade(uint(studentUserID), uint(teacherUserID), uint(subjectID), body.Value, body.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SchoolHandler) CreateAbsence(w http.ResponseWriter, r *http.Request) {
	studentUserIDStr := r.PathValue("userId")
	studentUserID, err := strconv.ParseUint(studentUserIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	subjectIDStr := r.PathValue("subjectId")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	var body struct {
		Date time.Time `json:"date"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateAbsence(uint(studentUserID), uint(subjectID), body.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SchoolHandler) GetMyClass(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	data, err := h.Service.GetMyClass(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *SchoolHandler) GetMyGradesAndAbsences(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	subjectIDStr := r.PathValue("subjectId")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	data, err := h.Service.GetMyGradesAndAbsences(uint(userID), uint(subjectID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *SchoolHandler) UpdateAbsenceStatus(w http.ResponseWriter, r *http.Request) {

	absenceIDStr := r.PathValue("id")
	absenceID, err := strconv.ParseUint(absenceIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid absence ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdateAbsenceStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Status != models.Excused &&
		req.Status != models.Unexcused &&
		req.Status != models.Pending {

		http.Error(w, "Invalid absence status", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateAbsenceStatus(uint(absenceID), req.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SchoolHandler) GetMyAverageGrade(w http.ResponseWriter, r *http.Request) {

	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	subjectIDStr := r.PathValue("subjectId")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	response, err := h.Service.GetMyAverageGrade(uint(userID), uint(subjectID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SchoolHandler) GetMyAbsenceStats(w http.ResponseWriter, r *http.Request) {

	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	subjectIDStr := r.PathValue("subjectId")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	response, err := h.Service.GetMyAbsenceStats(uint(userID), uint(subjectID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SchoolHandler) GetAverageGrade(w http.ResponseWriter, r *http.Request) {

	studentUserIDStr := r.PathValue("userId")
	studentUserID, err := strconv.ParseUint(studentUserIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	subjectIDStr := r.PathValue("subjectId")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	response, err := h.Service.GetAverageGrade(uint(studentUserID), uint(subjectID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SchoolHandler) GetAbsenceStats(w http.ResponseWriter, r *http.Request) {

	studentUserIDStr := r.PathValue("userId")
	studentUserID, err := strconv.ParseUint(studentUserIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	subjectIDStr := r.PathValue("subjectId")
	subjectID, err := strconv.ParseUint(subjectIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	response, err := h.Service.GetAbsenceStats(uint(studentUserID), uint(subjectID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SchoolHandler) GetStudentOverview(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	data, err := h.Service.GetStudentOverview(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *SchoolHandler) GetStudentAbsenceOverview(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	data, err := h.Service.GetStudentAbsenceOverview(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *SchoolHandler) GetStudentSummary(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	data, err := h.Service.GetStudentSummary(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *SchoolHandler) UpdateMyConsultationHours(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var body struct {
		ConsultationHours string `json:"consultationHours"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateTeacherConsultationHours(uint(userID), body.ConsultationHours); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SchoolHandler) GetChatbotContext(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetChatbotContext()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
