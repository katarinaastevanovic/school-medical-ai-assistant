package repository

import (
	"school-service/database"
	"school-service/dto"
	"school-service/models"
)

type SchoolRepository interface {
	CreateClass(class *models.SchoolClass) error
	GetAllClasses() ([]models.SchoolClass, error)
	GetStudentsFromClass(classID uint) ([]models.Student, error)
	AddStudentToClass(student *models.Student) error
	RemoveStudentFromClass(studentID uint) error
	GetAssignedStudents() ([]uint, error)
	CreateSubject(subject *models.Subject) error
	GetAllSubjects() ([]models.Subject, error)
	UpdateSubject(subject *models.Subject) error
	DeleteSubject(id uint) error
	GetOrCreateTeacher(userID uint) (models.Teacher, error)
	AssignTeacher(tcs *models.TeacherClassSubject) error
	AddTeacherToSubject(teacherID, subjectID uint) error
	GetTeachersBySubject(subjectID uint) ([]models.Teacher, error)
	GetSubjectsByClass(classID uint) ([]models.TeacherClassSubject, error)
	GetSubjectsForTeacher(userID uint) ([]models.TeacherClassSubject, error)
	GetGradesForStudent(studentID, subjectID uint) ([]models.Grade, error)
	GetAbsencesForStudent(studentID, subjectID uint) ([]models.Absence, error)
	CreateGrade(grade *models.Grade) error
	CreateAbsence(absence *models.Absence) error
	GetStudentByUserID(userID uint) (models.Student, error)
	UpdateAbsenceStatus(absenceID uint, status models.AbsenceType) error
	GetAverageGrade(studentID, subjectID uint) (float64, error)
	GetAbsenceStats(studentID, subjectID uint) (*dto.AbsenceStatsResponse, error)
	GetStudentSummary(studentID uint) (*dto.StudentSummaryResponse, error)
	UpdateTeacherConsultationHours(userID uint, hours string) error
	GetAllTeachers() ([]models.Teacher, error)
}

type schoolRepository struct{}

func NewSchoolRepository() SchoolRepository {
	return &schoolRepository{}
}

func (r *schoolRepository) CreateClass(class *models.SchoolClass) error {
	return database.DB.Create(class).Error
}

func (r *schoolRepository) GetAllClasses() ([]models.SchoolClass, error) {
	var classes []models.SchoolClass
	result := database.DB.Find(&classes)
	return classes, result.Error
}

func (r *schoolRepository) GetStudentsFromClass(classID uint) ([]models.Student, error) {
	var students []models.Student
	result := database.DB.Where("class_id = ?", classID).Find(&students)
	return students, result.Error
}

func (r *schoolRepository) AddStudentToClass(student *models.Student) error {
	var existing models.Student

	err := database.DB.Unscoped().
		Where("user_id = ?", student.UserID).
		First(&existing).Error

	if err == nil {
		// red postoji (mozda soft deleted), updateuj ga
		return database.DB.Unscoped().Model(&existing).Updates(map[string]interface{}{
			"class_id":   student.ClassID,
			"deleted_at": nil,
		}).Error
	}

	return database.DB.Create(student).Error
}

func (r *schoolRepository) RemoveStudentFromClass(studentID uint) error {
	return database.DB.Delete(&models.Student{}, studentID).Error
}

func (r *schoolRepository) GetAssignedStudents() ([]uint, error) {
	var ids []uint
	result := database.DB.Model(&models.Student{}).Pluck("user_id", &ids)
	return ids, result.Error
}

func (r *schoolRepository) CreateSubject(subject *models.Subject) error {
	return database.DB.Create(subject).Error
}

func (r *schoolRepository) GetAllSubjects() ([]models.Subject, error) {
	var subjects []models.Subject
	result := database.DB.Find(&subjects)
	return subjects, result.Error
}

func (r *schoolRepository) UpdateSubject(subject *models.Subject) error {
	return database.DB.Save(subject).Error
}

func (r *schoolRepository) DeleteSubject(id uint) error {
	return database.DB.Delete(&models.Subject{}, id).Error
}

func (r *schoolRepository) GetOrCreateTeacher(userID uint) (models.Teacher, error) {
	var teacher models.Teacher

	err := database.DB.Where("user_id = ?", userID).First(&teacher).Error
	if err == nil {
		return teacher, nil
	}

	teacher = models.Teacher{UserID: userID}
	err = database.DB.Create(&teacher).Error
	return teacher, err
}

func (r *schoolRepository) AssignTeacher(tcs *models.TeacherClassSubject) error {
	return database.DB.Create(tcs).Error
}

func (r *schoolRepository) AddTeacherToSubject(teacherID, subjectID uint) error {
	teacher := models.Teacher{}
	teacher.ID = teacherID

	subject := models.Subject{}
	subject.ID = subjectID

	return database.DB.Model(&teacher).Association("Subjects").Append(&subject)
}

func (r *schoolRepository) GetTeachersBySubject(subjectID uint) ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := database.DB.
		Joins("JOIN teacher_subjects ON teacher_subjects.teacher_id = teachers.id").
		Where("teacher_subjects.subject_id = ? AND teachers.deleted_at IS NULL", subjectID).
		Find(&teachers).Error
	return teachers, err
}

func (r *schoolRepository) GetSubjectsByClass(classID uint) ([]models.TeacherClassSubject, error) {
	var tcs []models.TeacherClassSubject
	result := database.DB.
		Preload("Subject").
		Preload("Teacher").
		Where("class_id = ?", classID).
		Find(&tcs)
	return tcs, result.Error
}

func (r *schoolRepository) GetSubjectsForTeacher(userID uint) ([]models.TeacherClassSubject, error) {
	var teacher models.Teacher
	if err := database.DB.Where("user_id = ?", userID).First(&teacher).Error; err != nil {
		return nil, err
	}

	var tcs []models.TeacherClassSubject
	err := database.DB.
		Preload("Subject").
		Preload("Class").
		Where("teacher_id = ?", teacher.ID).
		Find(&tcs).Error

	return tcs, err
}

func (r *schoolRepository) GetGradesForStudent(studentID, subjectID uint) ([]models.Grade, error) {
	var grades []models.Grade
	err := database.DB.
		Where("student_id = ? AND subject_id = ?", studentID, subjectID).
		Order("date desc").
		Find(&grades).Error
	return grades, err
}

func (r *schoolRepository) GetAbsencesForStudent(studentID, subjectID uint) ([]models.Absence, error) {
	var absences []models.Absence
	err := database.DB.
		Where("student_id = ? AND subject_id = ?", studentID, subjectID).
		Order("date desc").
		Find(&absences).Error
	return absences, err
}

func (r *schoolRepository) CreateGrade(grade *models.Grade) error {
	return database.DB.Create(grade).Error
}

func (r *schoolRepository) CreateAbsence(absence *models.Absence) error {
	return database.DB.Create(absence).Error
}

func (r *schoolRepository) GetStudentByUserID(userID uint) (models.Student, error) {
	var student models.Student
	err := database.DB.Where("user_id = ?", userID).First(&student).Error
	return student, err
}

func (r *schoolRepository) UpdateAbsenceStatus(absenceID uint, status models.AbsenceType) error {
	return database.DB.Model(&models.Absence{}).
		Where("id = ?", absenceID).
		Update("status", status).Error
}

func (r *schoolRepository) GetAverageGrade(studentID, subjectID uint) (float64, error) {
	var average float64

	err := database.DB.
		Model(&models.Grade{}).
		Where("student_id = ? AND subject_id = ?", studentID, subjectID).
		Select("COALESCE(AVG(value), 0)"). //umesto null uzme 0
		Scan(&average).Error

	return average, err
}

func (r *schoolRepository) GetAbsenceStats(studentID, subjectID uint) (*dto.AbsenceStatsResponse, error) {

	stats := &dto.AbsenceStatsResponse{}

	err := database.DB.
		Model(&models.Absence{}).
		Where("student_id = ? AND subject_id = ? AND status = ?", studentID, subjectID, models.Pending).
		Count(&stats.Pending).Error
	if err != nil {
		return nil, err
	}

	err = database.DB.
		Model(&models.Absence{}).
		Where("student_id = ? AND subject_id = ? AND status = ?", studentID, subjectID, models.Excused).
		Count(&stats.Excused).Error
	if err != nil {
		return nil, err
	}

	err = database.DB.
		Model(&models.Absence{}).
		Where("student_id = ? AND subject_id = ? AND status = ?", studentID, subjectID, models.Unexcused).
		Count(&stats.Unexcused).Error
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *schoolRepository) GetStudentSummary(studentID uint) (*dto.StudentSummaryResponse, error) {
	var result dto.StudentSummaryResponse

	err := database.DB.
		Model(&models.Absence{}).
		Where("student_id = ? AND status = ?", studentID, models.Unexcused).
		Count(&result.Unexcused).Error

	if err != nil {
		return nil, err
	}

	err = database.DB.
		Model(&models.Absence{}).
		Where("student_id = ? AND status = ?", studentID, models.Excused).
		Count(&result.Excused).Error

	if err != nil {
		return nil, err
	}

	err = database.DB.
		Model(&models.Absence{}).
		Where("student_id = ? AND status = ?", studentID, models.Pending).
		Count(&result.Pending).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *schoolRepository) UpdateTeacherConsultationHours(userID uint, hours string) error {
	return database.DB.Model(&models.Teacher{}).
		Where("user_id = ?", userID).
		Update("consultation_hours", hours).Error
}

func (r *schoolRepository) GetAllTeachers() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := database.DB.Find(&teachers).Error
	return teachers, err
}
