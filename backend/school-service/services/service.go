package services

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
	"school-service/database"
	"school-service/dto"
	"school-service/models"
	"school-service/repository"
	"time"
)

type SchoolService struct {
	Repo repository.SchoolRepository
}

func (s *SchoolService) CreateClass(class *models.SchoolClass) error {

	if class.Name == "" {
		return errors.New("class name is required")
	}

	if class.Year <= 0 {
		return errors.New("invalid year")
	}

	return s.Repo.CreateClass(class)
}

func (s *SchoolService) GetAllClasses() ([]models.SchoolClass, error) {
	return s.Repo.GetAllClasses()
}

func (s *SchoolService) GetStudentsFromClass(classID uint) ([]models.Student, error) {
	return s.Repo.GetStudentsFromClass(classID)
}

func (s *SchoolService) AddStudentToClass(student *models.Student) error {
	if student.UserID == 0 {
		return errors.New("user ID is required")
	}

	if student.ClassID == 0 {
		return errors.New("class ID is required")
	}

	return s.Repo.AddStudentToClass(student)
}

func (s *SchoolService) RemoveStudentFromClass(studentID uint) error {
	if studentID == 0 {
		return errors.New("student ID is required")
	}

	return s.Repo.RemoveStudentFromClass(studentID)
}

func (s *SchoolService) GetFreeStudents() ([]map[string]interface{}, error) {
	assignedIDs, err := s.Repo.GetAssignedStudents()
	if err != nil {
		return nil, err
	}

	assignedSet := make(map[uint]bool)
	for _, id := range assignedIDs {
		assignedSet[id] = true
	}

	resp, err := http.Get("http://auth-service:8081/members/students")
	if err != nil {
		return nil, errors.New("failed to reach auth service")
	}
	defer resp.Body.Close()

	var allStudents []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&allStudents); err != nil {
		return nil, errors.New("failed to parse auth service response")
	}

	free := make([]map[string]interface{}, 0)
	for _, student := range allStudents {
		idFloat, ok := student["id"].(float64)
		if !ok {
			continue
		}
		if !assignedSet[uint(idFloat)] {
			free = append(free, student)
		}
	}

	log.Println("Free students:", free)
	return free, nil
}

func (s *SchoolService) CreateSubject(subject *models.Subject) error {
	if subject.Name == "" {
		return errors.New("subject name is required")
	}
	return s.Repo.CreateSubject(subject)
}

func (s *SchoolService) GetAllSubjects() ([]models.Subject, error) {
	return s.Repo.GetAllSubjects()
}

func (s *SchoolService) UpdateSubject(id uint, name string) error {
	if name == "" {
		return errors.New("subject name is required")
	}
	subject := &models.Subject{}
	subject.ID = id
	subject.Name = name
	return s.Repo.UpdateSubject(subject)
}

func (s *SchoolService) DeleteSubject(id uint) error {
	if id == 0 {
		return errors.New("subject ID is required")
	}
	return s.Repo.DeleteSubject(id)
}

func (s *SchoolService) AssignTeacher(userID, classID, subjectID uint) error {
	if userID == 0 {
		return errors.New("user ID is required")
	}
	if classID == 0 {
		return errors.New("class ID is required")
	}
	if subjectID == 0 {
		return errors.New("subject ID is required")
	}

	teacher, err := s.Repo.GetOrCreateTeacher(userID)
	if err != nil {
		return err
	}

	tcs := &models.TeacherClassSubject{
		TeacherID: teacher.ID,
		ClassID:   classID,
		SubjectID: subjectID,
	}

	return s.Repo.AssignTeacher(tcs)
}

func (s *SchoolService) AddTeacherToSubject(userID, subjectID uint) error {
	if userID == 0 {
		return errors.New("user ID is required")
	}
	if subjectID == 0 {
		return errors.New("subject ID is required")
	}

	teacher, err := s.Repo.GetOrCreateTeacher(userID)
	if err != nil {
		return err
	}

	return s.Repo.AddTeacherToSubject(teacher.ID, subjectID)
}

func (s *SchoolService) GetTeachersBySubjectEnriched(subjectID uint) ([]map[string]interface{}, error) {
	teachers, err := s.Repo.GetTeachersBySubject(subjectID)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("http://auth-service:8081/members/teachers")
	if err != nil {
		return nil, errors.New("failed to reach auth service")
	}
	defer resp.Body.Close()

	var allTeachers []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&allTeachers); err != nil {
		return nil, errors.New("failed to parse auth service response")
	}

	authMap := make(map[uint]map[string]interface{})
	for _, t := range allTeachers {
		idFloat, ok := t["id"].(float64)
		if !ok {
			continue
		}
		authMap[uint(idFloat)] = t
	}

	result := make([]map[string]interface{}, 0)
	for _, t := range teachers {
		authData, ok := authMap[t.UserID]
		if !ok {
			continue
		}
		result = append(result, map[string]interface{}{
			"id":                t.ID,
			"userID":            t.UserID,
			"firstName":         authData["firstName"],
			"lastName":          authData["lastName"],
			"consultationHours": t.ConsultationHours,
		})
	}

	return result, nil
}

func (s *SchoolService) GetAllTeachersEnriched() ([]map[string]interface{}, error) {
	teachers, err := s.Repo.GetAllTeachers()
	if err != nil {
		return nil, err
	}

	hoursMap := make(map[uint]string)
	for _, t := range teachers {
		hoursMap[t.UserID] = t.ConsultationHours
	}

	resp, err := http.Get("http://auth-service:8081/members/teachers")
	if err != nil {
		return nil, errors.New("failed to reach auth service")
	}
	defer resp.Body.Close()

	var allTeachers []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&allTeachers); err != nil {
		return nil, errors.New("failed to parse auth service response")
	}

	result := make([]map[string]interface{}, 0)
	for _, t := range allTeachers {
		idFloat, ok := t["id"].(float64)
		if !ok {
			continue
		}
		userID := uint(idFloat)
		result = append(result, map[string]interface{}{
			"userID":            userID,
			"firstName":         t["firstName"],
			"lastName":          t["lastName"],
			"consultationHours": hoursMap[userID],
		})
	}

	return result, nil
}

func (s *SchoolService) GetSubjectsByClass(classID uint) ([]map[string]interface{}, error) {
	tcs, err := s.Repo.GetSubjectsByClass(classID)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("http://auth-service:8081/members/teachers")
	if err != nil {
		return nil, errors.New("failed to reach auth service")
	}
	defer resp.Body.Close()

	var allTeachers []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&allTeachers); err != nil {
		return nil, errors.New("failed to parse auth service response")
	}

	authMap := make(map[uint]map[string]interface{})
	for _, t := range allTeachers {
		idFloat, ok := t["id"].(float64)
		if !ok {
			continue
		}
		authMap[uint(idFloat)] = t
	}

	result := make([]map[string]interface{}, 0)
	for _, tc := range tcs {
		authData, ok := authMap[tc.Teacher.UserID]
		if !ok {
			continue
		}
		result = append(result, map[string]interface{}{
			"id":          tc.ID,
			"subjectID":   tc.SubjectID,
			"subjectName": tc.Subject.Name,
			"teacherID":   tc.TeacherID,
			"userID":      tc.Teacher.UserID,
			"firstName":   authData["firstName"],
			"lastName":    authData["lastName"],
		})
	}

	return result, nil
}

func (s *SchoolService) GetSubjectsForTeacher(userID uint) ([]map[string]interface{}, error) {
	tcs, err := s.Repo.GetSubjectsForTeacher(userID)
	if err != nil {
		return make([]map[string]interface{}, 0), nil
	}

	result := make([]map[string]interface{}, 0)
	for _, tc := range tcs {
		result = append(result, map[string]interface{}{
			"id":          tc.ID,
			"subjectID":   tc.SubjectID,
			"subjectName": tc.Subject.Name,
			"classID":     tc.ClassID,
			"className":   tc.Class.Name,
		})
	}

	return result, nil
}

func (s *SchoolService) GetGradesAndAbsences(studentUserID, subjectID uint) (*dto.GradesAndAbsencesResponse, error) {
	student, err := s.Repo.GetStudentByUserID(studentUserID)
	if err != nil {
		return nil, errors.New("student not found")
	}

	grades, err := s.Repo.GetGradesForStudent(student.ID, subjectID)
	if err != nil {
		return nil, err
	}

	absences, err := s.Repo.GetAbsencesForStudent(student.ID, subjectID)
	if err != nil {
		return nil, err
	}

	gradeResponses := make([]dto.GradeResponse, 0, len(grades))
	for _, g := range grades {
		gradeResponses = append(gradeResponses, dto.GradeResponse{
			ID:        g.ID,
			Value:     g.Value,
			Date:      g.Date,
			StudentID: g.StudentID,
			TeacherID: g.TeacherID,
			SubjectID: g.SubjectID,
		})
	}

	absenceResponses := make([]dto.AbsenceResponse, 0, len(absences))
	for _, a := range absences {
		absenceResponses = append(absenceResponses, dto.AbsenceResponse{
			ID:        a.ID,
			Status:    string(a.Status),
			Date:      a.Date,
			StudentID: a.StudentID,
			SubjectID: a.SubjectID,
		})
	}

	return &dto.GradesAndAbsencesResponse{
		Grades:   gradeResponses,
		Absences: absenceResponses,
	}, nil
}

func (s *SchoolService) CreateGrade(studentUserID, teacherUserID, subjectID uint, value int, date time.Time) error {
	if value < 1 || value > 5 {
		return errors.New("grade value must be between 1 and 5")
	}

	student, err := s.Repo.GetStudentByUserID(studentUserID)
	if err != nil {
		return errors.New("student not found")
	}

	teacher, err := s.Repo.GetOrCreateTeacher(teacherUserID)
	if err != nil {
		return err
	}

	grade := &models.Grade{
		Value:     value,
		Date:      date,
		StudentID: student.ID,
		TeacherID: teacher.ID,
		SubjectID: subjectID,
	}

	return s.Repo.CreateGrade(grade)
}

func (s *SchoolService) CreateAbsence(studentUserID, subjectID uint, date time.Time) error {
	student, err := s.Repo.GetStudentByUserID(studentUserID)
	if err != nil {
		return errors.New("student not found")
	}

	absence := &models.Absence{
		StudentID: student.ID,
		SubjectID: subjectID,
		Status:    models.Pending,
		Date:      date,
	}

	return s.Repo.CreateAbsence(absence)
}

func (s *SchoolService) GetMyClass(userID uint) (map[string]interface{}, error) {
	student, err := s.Repo.GetStudentByUserID(userID)
	if err != nil {
		return nil, errors.New("student not found or not assigned to a class")
	}

	var class models.SchoolClass
	if err := database.DB.First(&class, student.ClassID).Error; err != nil {
		return nil, err
	}

	subjects, err := s.GetSubjectsByClass(student.ClassID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"classID":   student.ClassID,
		"className": class.Name,
		"subjects":  subjects,
	}, nil
}

func (s *SchoolService) GetMyGradesAndAbsences(userID, subjectID uint) (*dto.GradesAndAbsencesResponse, error) {
	return s.GetGradesAndAbsences(userID, subjectID)
}

func (s *SchoolService) UpdateAbsenceStatus(absenceID uint, status models.AbsenceType) error {
	return s.Repo.UpdateAbsenceStatus(absenceID, status)
}

func (s *SchoolService) GetMyAverageGrade(userID, subjectID uint) (*dto.AverageGradeResponse, error) {

	student, err := s.Repo.GetStudentByUserID(userID)
	if err != nil {
		return nil, err
	}

	average, err := s.Repo.GetAverageGrade(student.ID, subjectID)
	if err != nil {
		return nil, err
	}

	average = math.Round(average*100) / 100

	return &dto.AverageGradeResponse{
		Average: average,
	}, nil
}

func (s *SchoolService) GetMyAbsenceStats(userID, subjectID uint) (*dto.AbsenceStatsResponse, error) {

	student, err := s.Repo.GetStudentByUserID(userID)
	if err != nil {
		return nil, err
	}

	return s.Repo.GetAbsenceStats(student.ID, subjectID)
}

func (s *SchoolService) GetAverageGrade(userID, subjectID uint) (*dto.AverageGradeResponse, error) {

	student, err := s.Repo.GetStudentByUserID(userID)
	if err != nil {
		return nil, err
	}

	average, err := s.Repo.GetAverageGrade(student.ID, subjectID)
	if err != nil {
		return nil, err
	}

	average = math.Round(average*100) / 100

	return &dto.AverageGradeResponse{
		Average: average,
	}, nil
}

func (s *SchoolService) GetAbsenceStats(userID, subjectID uint) (*dto.AbsenceStatsResponse, error) {

	student, err := s.Repo.GetStudentByUserID(userID)
	if err != nil {
		return nil, err
	}

	return s.Repo.GetAbsenceStats(student.ID, subjectID)
}

func (s *SchoolService) GetStudentOverview(userID uint) ([]map[string]interface{}, error) {
	student, err := s.Repo.GetStudentByUserID(userID)
	if err != nil {
		return nil, errors.New("student not found or not assigned to a class")
	}

	tcs, err := s.Repo.GetSubjectsByClass(student.ClassID)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)
	for _, tc := range tcs {
		avg, err := s.Repo.GetAverageGrade(student.ID, tc.SubjectID)
		if err != nil {
			avg = 0
		}
		avg = math.Round(avg*100) / 100

		result = append(result, map[string]interface{}{
			"subjectID":   tc.SubjectID,
			"subjectName": tc.Subject.Name,
			"average":     avg,
		})
	}

	return result, nil
}

func (s *SchoolService) GetStudentAbsenceOverview(userID uint) ([]map[string]interface{}, error) {
	student, err := s.Repo.GetStudentByUserID(userID)
	if err != nil {
		return nil, errors.New("student not found or not assigned to a class")
	}

	tcs, err := s.Repo.GetSubjectsByClass(student.ClassID)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)
	for _, tc := range tcs {
		stats, err := s.Repo.GetAbsenceStats(student.ID, tc.SubjectID)
		if err != nil {
			continue
		}
		result = append(result, map[string]interface{}{
			"subjectID":   tc.SubjectID,
			"subjectName": tc.Subject.Name,
			"pending":     stats.Pending,
			"excused":     stats.Excused,
			"unexcused":   stats.Unexcused,
		})
	}

	return result, nil
}

func (s *SchoolService) GetStudentSummary(userID uint) (*dto.StudentSummaryResponse, error) {
	student, err := s.Repo.GetStudentByUserID(userID)
	if err != nil {
		return nil, errors.New("student not found")
	}

	summary, err := s.Repo.GetStudentSummary(student.ID)
	if err != nil {
		return nil, err
	}

	// Uzimamo sve predmete koje student ima
	tcs, err := s.Repo.GetSubjectsByClass(student.ClassID)
	if err != nil {
		return nil, err
	}

	var totalAverage float64
	var subjectCount int

	for _, tc := range tcs {
		average, err := s.Repo.GetAverageGrade(student.ID, tc.SubjectID)
		if err != nil {
			return nil, err
		}

		// Predmet bez ocena ne ulazi u prosek
		if average > 0 {
			totalAverage += average
			subjectCount++
		}
	}

	// Prosek proseka svih predmeta
	if subjectCount > 0 {
		summary.AverageGrade = totalAverage / float64(subjectCount)
	} else {
		summary.AverageGrade = 0
	}

	summary.AverageGrade = math.Round(summary.AverageGrade*100) / 100

	return summary, nil
}

func (s *SchoolService) UpdateTeacherConsultationHours(userID uint, hours string) error {
	return s.Repo.UpdateTeacherConsultationHours(userID, hours)
}

func (s *SchoolService) GetChatbotContext() (map[string]interface{}, error) {
	subjects, err := s.Repo.GetAllSubjects()
	if err != nil {
		return nil, err
	}

	classes, err := s.Repo.GetAllClasses()
	if err != nil {
		return nil, err
	}
	classNames := make([]string, 0, len(classes))
	for _, c := range classes {
		classNames = append(classNames, c.Name)
	}

	subjectsWithTeachers := make([]map[string]interface{}, 0, len(subjects))
	for _, sub := range subjects {
		teachers, err := s.GetTeachersBySubjectEnriched(sub.ID)
		if err != nil {
			teachers = []map[string]interface{}{}
		}

		teacherInfo := make([]map[string]interface{}, 0, len(teachers))
		for _, t := range teachers {
			teacherInfo = append(teacherInfo, map[string]interface{}{
				"firstName":         t["firstName"],
				"lastName":          t["lastName"],
				"consultationHours": t["consultationHours"],
			})
		}

		subjectsWithTeachers = append(subjectsWithTeachers, map[string]interface{}{
			"subjectName": sub.Name,
			"teachers":    teacherInfo,
		})
	}

	return map[string]interface{}{
		"classes":  classNames,
		"subjects": subjectsWithTeachers,
	}, nil
}
