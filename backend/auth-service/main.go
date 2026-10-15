package main

import (
	"log"
	"net/http"

	"auth-service/database"
	"auth-service/handlers"
	"auth-service/repository"
	"auth-service/services"

	"github.com/joho/godotenv"
)

func main() {

	// load env
	err := godotenv.Load()

	// database
	database.Connect()

	// repository
	authRepo := repository.NewAuthRepository()

	// services
	authService := services.AuthService{
		Repo: authRepo,
	}

	// handlers
	authHandler := handlers.AuthHandler{
		Service: authService,
	}

	// mux router
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/register", authHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)

	mux.HandleFunc("/members", authHandler.GetAllMembers)
	mux.HandleFunc("GET /members/students", authHandler.GetAllStudents)

	mux.HandleFunc("GET /members/{id}", authHandler.GetMemberByID)
	mux.HandleFunc("PUT /members/{id}", authHandler.UpdateMember)

	mux.HandleFunc("GET /members/teachers", authHandler.GetAllTeachers)
	mux.HandleFunc("GET /members/doctors", authHandler.GetAllDoctors)

	log.Println("Auth service running on :8081")

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
