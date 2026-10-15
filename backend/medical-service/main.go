package main

import (
	"log"
	"net/http"

	"medical-service/database"
	"medical-service/handlers"
	"medical-service/repository"
	"medical-service/services"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.Connect()

	medicalRepo := repository.NewMedicalRepository()
	medicalService := services.MedicalService{Repo: medicalRepo}
	medicalHandler := handlers.MedicalHandler{Service: medicalService}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /examination-requests", medicalHandler.CreateExaminationRequest)
	mux.HandleFunc("GET /examination-requests", medicalHandler.GetExaminationRequests)
	mux.HandleFunc("GET /examination-requests/me", medicalHandler.GetMyExaminationRequests)
	mux.HandleFunc("PUT /examination-requests/{id}/status", medicalHandler.UpdateRequestStatus)
	mux.HandleFunc("POST /examination-requests/{id}/examination", medicalHandler.CreateExamination)
	mux.HandleFunc("POST /certificates", medicalHandler.CreateCertificate)
	mux.HandleFunc("GET /certificates/me", medicalHandler.GetMyCertificates)
	mux.HandleFunc("GET /medical-record/me", medicalHandler.GetMyMedicalRecord)
	mux.HandleFunc("GET /patients", medicalHandler.GetAllPatients)
	mux.HandleFunc("GET /certificates/{userID}", medicalHandler.GetCertificatesByUserID)
	mux.HandleFunc("GET /medical-record/{userID}", medicalHandler.GetMedicalRecordByUserID)
	mux.HandleFunc("PUT /doctors/me/working-hours", medicalHandler.UpdateMyWorkingHours)
	mux.HandleFunc("GET /doctors", medicalHandler.GetAllDoctorsEnriched)
	mux.HandleFunc("GET /internal/chatbot-context", medicalHandler.GetChatbotContext)

	_ = medicalHandler

	log.Println("Medical service running on :8083")
	if err := http.ListenAndServe(":8083", mux); err != nil {
		log.Fatal(err)
	}
}
