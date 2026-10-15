package main

import (
	"log"
	"net/http"
	"os"

	"gateway-service/middleware"
	"gateway-service/proxy"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load()

	authURL := os.Getenv("AUTH_SERVICE_URL")
	schoolURL := os.Getenv("SCHOOL_SERVICE_URL")
	medicalURL := os.Getenv("MEDICAL_SERVICE_URL")

	authProxy := proxy.NewProxy(authURL)
	schoolProxy := proxy.NewProxy(schoolURL)
	medicalProxy := proxy.NewProxy(medicalURL)

	mux := http.NewServeMux()

	// public rute
	mux.Handle("POST /login", authProxy)
	mux.Handle("POST /register", authProxy)

	// auth servis
	mux.Handle("GET /members/teachers", middleware.Auth(middleware.RequireRole("ADMIN")(authProxy)))
	mux.Handle("GET /members/students", middleware.Auth(middleware.RequireRole("ADMIN")(authProxy)))
	mux.Handle("GET /members/{id}", middleware.Auth(authProxy))
	mux.Handle("GET /members", middleware.Auth(middleware.RequireRole("ADMIN")(authProxy)))
	mux.Handle("PUT /members/{id}", middleware.Auth(middleware.RequireRole("ADMIN", "TEACHER", "STUDENT", "DOCTOR")(authProxy)))

	// school servis
	mux.Handle("POST /class", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))
	mux.Handle("GET /classes", middleware.Auth(middleware.RequireRole("ADMIN", "TEACHER", "STUDENT")(schoolProxy)))
	mux.Handle("GET /classes/{id}/students", middleware.Auth(middleware.RequireRole("ADMIN", "TEACHER")(schoolProxy)))
	mux.Handle("POST /classes/{id}/students/add", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))
	mux.Handle("DELETE /classes/{id}/students/{studentId}", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))
	mux.Handle("GET /students/free", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))

	mux.Handle("POST /subjects", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))
	mux.Handle("GET /subjects", middleware.Auth(middleware.RequireRole("ADMIN", "TEACHER", "STUDENT")(schoolProxy)))
	mux.Handle("PUT /subjects/{id}", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))
	mux.Handle("DELETE /subjects/{id}", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))

	mux.Handle("POST /teachers/assign", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))
	mux.Handle("POST /subjects/{id}/teachers", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))
	mux.Handle("GET /subjects/{id}/teachers", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))
	mux.Handle("GET /teachers", middleware.Auth(middleware.RequireRole("ADMIN")(schoolProxy)))
	mux.Handle("GET /teachers/me/subjects", middleware.Auth(middleware.RequireRole("TEACHER")(schoolProxy)))

	mux.Handle("GET /classes/{id}/subjects", middleware.Auth(middleware.RequireRole("ADMIN", "TEACHER", "STUDENT")(schoolProxy)))

	mux.Handle("GET /students/{userId}/subjects/{subjectId}/grades-absences", middleware.Auth(middleware.RequireRole("ADMIN", "TEACHER", "STUDENT")(schoolProxy)))
	mux.Handle("POST /students/{userId}/subjects/{subjectId}/grades", middleware.Auth(middleware.RequireRole("TEACHER")(schoolProxy)))
	mux.Handle("POST /students/{userId}/subjects/{subjectId}/absences", middleware.Auth(middleware.RequireRole("TEACHER")(schoolProxy)))
	mux.Handle("GET /students/me/class", middleware.Auth(middleware.RequireRole("STUDENT")(schoolProxy)))
	mux.Handle("GET /students/me/subjects/{subjectId}/grades-absences", middleware.Auth(middleware.RequireRole("STUDENT")(schoolProxy)))
	mux.Handle("PUT /absences/{id}/status", middleware.Auth(middleware.RequireRole("TEACHER")(schoolProxy)))
	mux.Handle("GET /students/me/subjects/{subjectId}/average-grade", middleware.Auth(middleware.RequireRole("STUDENT", "TEACHER")(schoolProxy)))
	mux.Handle("GET /students/me/subjects/{subjectId}/absence-stats", middleware.Auth(middleware.RequireRole("STUDENT", "TEACHER")(schoolProxy)))
	mux.Handle("GET /students/{userId}/subjects/{subjectId}/average-grade", middleware.Auth(middleware.RequireRole("ADMIN", "TEACHER")(schoolProxy)))
	mux.Handle("GET /students/{userId}/subjects/{subjectId}/absence-stats", middleware.Auth(middleware.RequireRole("ADMIN", "TEACHER")(schoolProxy)))
	mux.Handle("GET /students/me/overview", middleware.Auth(middleware.RequireRole("STUDENT")(schoolProxy)))
	mux.Handle("GET /students/me/absence-overview", middleware.Auth(middleware.RequireRole("STUDENT")(schoolProxy)))
	mux.Handle("GET /students/me/summary", middleware.Auth(middleware.RequireRole("STUDENT")(schoolProxy)))
	mux.Handle("PUT /teachers/me/consultation-hours", middleware.Auth(middleware.RequireRole("TEACHER")(schoolProxy)))

	// medical servis
	mux.Handle("POST /examination-requests", middleware.Auth(middleware.RequireRole("STUDENT")(medicalProxy)))
	mux.Handle("GET /examination-requests", middleware.Auth(middleware.RequireRole("DOCTOR")(medicalProxy)))
	mux.Handle("GET /examination-requests/me", middleware.Auth(middleware.RequireRole("STUDENT")(medicalProxy)))
	mux.Handle("PUT /examination-requests/{id}/status", middleware.Auth(middleware.RequireRole("DOCTOR")(medicalProxy)))
	mux.Handle("POST /examination-requests/{id}/examination", middleware.Auth(middleware.RequireRole("DOCTOR")(medicalProxy)))
	mux.Handle("POST /certificates", middleware.Auth(middleware.RequireRole("DOCTOR")(medicalProxy)))
	mux.Handle("GET /certificates/me", middleware.Auth(middleware.RequireRole("STUDENT")(medicalProxy)))
	mux.Handle("GET /medical-record/me", middleware.Auth(middleware.RequireRole("STUDENT")(medicalProxy)))
	mux.Handle("GET /patients", middleware.Auth(middleware.RequireRole("DOCTOR")(medicalProxy)))
	mux.Handle("GET /members/doctors", middleware.Auth(middleware.RequireRole("STUDENT", "ADMIN")(authProxy)))
	mux.Handle("GET /certificates/{userID}", middleware.Auth(middleware.RequireRole("TEACHER", "ADMIN", "DOCTOR")(medicalProxy)))
	mux.Handle("GET /medical-record/{userID}", middleware.Auth(middleware.RequireRole("DOCTOR", "ADMIN")(medicalProxy)))
	mux.Handle("PUT /doctors/me/working-hours", middleware.Auth(middleware.RequireRole("DOCTOR")(medicalProxy)))
	mux.Handle("GET /doctors", middleware.Auth(middleware.RequireRole("STUDENT", "ADMIN", "DOCTOR")(medicalProxy)))

	//chatbot
	chatbotURL := os.Getenv("CHATBOT_SERVICE_URL")
	chatbotProxy := proxy.NewProxy(chatbotURL)

	mux.Handle("POST /chatbot/ask", middleware.Auth(middleware.RequireRole("ADMIN", "TEACHER", "STUDENT", "DOCTOR")(chatbotProxy)))
	mux.Handle("GET /chatbot/instructions", middleware.Auth(middleware.RequireRole("ADMIN")(chatbotProxy)))
	mux.Handle("POST /chatbot/instructions", middleware.Auth(middleware.RequireRole("ADMIN")(chatbotProxy)))
	mux.Handle("PUT /chatbot/instructions/{id}", middleware.Auth(middleware.RequireRole("ADMIN")(chatbotProxy)))
	mux.Handle("DELETE /chatbot/instructions/{id}", middleware.Auth(middleware.RequireRole("ADMIN")(chatbotProxy)))
	mux.Handle("GET /chatbot/settings", middleware.Auth(middleware.RequireRole("ADMIN")(chatbotProxy)))
	mux.Handle("PUT /chatbot/settings", middleware.Auth(middleware.RequireRole("ADMIN")(chatbotProxy)))

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := corsMiddleware.Handler(mux)

	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
