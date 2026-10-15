package main

import (
	"log"
	"net/http"

	"school-service/database"
	"school-service/handlers"
	"school-service/repository"
	"school-service/services"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	database.Connect()

	schoolRepo := repository.NewSchoolRepository()

	schoolService := services.SchoolService{
		Repo: schoolRepo,
	}

	schoolHandler := handlers.SchoolHandler{
		Service: schoolService,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /class", schoolHandler.CreateClass)
	mux.HandleFunc("GET /classes", schoolHandler.GetAllClasses)
	mux.HandleFunc("GET /classes/{id}/students", schoolHandler.GetStudentsFromClass)
	mux.HandleFunc("POST /classes/{id}/students/add", schoolHandler.AddStudentToClass)
	mux.HandleFunc("DELETE /classes/{id}/students/{studentId}", schoolHandler.RemoveStudentFromClass)
	mux.HandleFunc("GET /students/free", schoolHandler.GetFreeStudents)
	mux.HandleFunc("POST /subjects", schoolHandler.CreateSubject)
	mux.HandleFunc("GET /subjects", schoolHandler.GetAllSubjects)
	mux.HandleFunc("PUT /subjects/{id}", schoolHandler.UpdateSubject)
	mux.HandleFunc("DELETE /subjects/{id}", schoolHandler.DeleteSubject)
	mux.HandleFunc("POST /teachers/assign", schoolHandler.AssignTeacher)
	mux.HandleFunc("POST /subjects/{id}/teachers", schoolHandler.AddTeacherToSubject)
	mux.HandleFunc("GET /subjects/{id}/teachers", schoolHandler.GetTeachersBySubjectEnriched)
	mux.HandleFunc("GET /teachers", schoolHandler.GetAllTeachersEnriched)
	mux.HandleFunc("GET /classes/{id}/subjects", schoolHandler.GetSubjectsByClass)
	mux.HandleFunc("GET /teachers/me/subjects", schoolHandler.GetMySubjects)
	mux.HandleFunc("GET /students/{userId}/subjects/{subjectId}/grades-absences", schoolHandler.GetGradesAndAbsences)
	mux.HandleFunc("POST /students/{userId}/subjects/{subjectId}/grades", schoolHandler.CreateGrade)
	mux.HandleFunc("POST /students/{userId}/subjects/{subjectId}/absences", schoolHandler.CreateAbsence)
	mux.HandleFunc("GET /students/me/class", schoolHandler.GetMyClass)
	mux.HandleFunc("GET /students/me/subjects/{subjectId}/grades-absences", schoolHandler.GetMyGradesAndAbsences)
	mux.HandleFunc("GET /students/{userId}/subjects/{subjectId}/average-grade", schoolHandler.GetAverageGrade)
	mux.HandleFunc("GET /students/{userId}/subjects/{subjectId}/absence-stats", schoolHandler.GetAbsenceStats)
	mux.HandleFunc("PUT /absences/{id}/status", schoolHandler.UpdateAbsenceStatus)
	mux.HandleFunc("GET /students/me/subjects/{subjectId}/average-grade", schoolHandler.GetMyAverageGrade)
	mux.HandleFunc("GET /students/me/subjects/{subjectId}/absence-stats", schoolHandler.GetMyAbsenceStats)
	mux.HandleFunc("GET /students/me/overview", schoolHandler.GetStudentOverview)
	mux.HandleFunc("GET /students/me/absence-overview", schoolHandler.GetStudentAbsenceOverview)
	mux.HandleFunc("GET /students/me/summary", schoolHandler.GetStudentSummary)
	mux.HandleFunc("PUT /teachers/me/consultation-hours", schoolHandler.UpdateMyConsultationHours)
	mux.HandleFunc("GET /internal/chatbot-context", schoolHandler.GetChatbotContext)

	log.Println("School service running on :8082")

	err = http.ListenAndServe(":8082", mux)
	if err != nil {
		log.Fatal(err)
	}
}
