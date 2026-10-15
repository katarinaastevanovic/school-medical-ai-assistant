package main

import (
	"log"
	"net/http"

	"chatbot-service/database"
	"chatbot-service/handlers"
	"chatbot-service/repository"
	"chatbot-service/services"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	repo := repository.NewChatbotRepository()
	service := services.ChatbotService{Repo: repo}
	handler := handlers.ChatbotHandler{Service: service}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /chatbot/ask", handler.Ask)
	mux.HandleFunc("GET /chatbot/instructions", handler.GetInstructions)
	mux.HandleFunc("POST /chatbot/instructions", handler.CreateInstruction)
	mux.HandleFunc("PUT /chatbot/instructions/{id}", handler.UpdateInstruction)
	mux.HandleFunc("DELETE /chatbot/instructions/{id}", handler.DeleteInstruction)
	mux.HandleFunc("GET /chatbot/settings", handler.GetSettings)
	mux.HandleFunc("PUT /chatbot/settings", handler.UpdateSettings)

	log.Println("Chatbot service running on :8084")
	log.Fatal(http.ListenAndServe(":8084", mux))
}
