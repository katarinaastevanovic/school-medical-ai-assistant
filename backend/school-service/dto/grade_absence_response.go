package dto

import "time"

type GradeResponse struct {
	ID        uint      `json:"id"`
	Value     int       `json:"value"`
	Date      time.Time `json:"date"`
	StudentID uint      `json:"studentId"`
	TeacherID uint      `json:"teacherId"`
	SubjectID uint      `json:"subjectId"`
}

type AbsenceResponse struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	Date      time.Time `json:"date"`
	StudentID uint      `json:"studentId"`
	SubjectID uint      `json:"subjectId"`
}

type GradesAndAbsencesResponse struct {
	Grades   []GradeResponse   `json:"grades"`
	Absences []AbsenceResponse `json:"absences"`
}
