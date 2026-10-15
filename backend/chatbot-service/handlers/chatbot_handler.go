package handlers

import (
	"chatbot-service/models"
	"chatbot-service/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type ChatbotHandler struct {
	Service services.ChatbotService
}

func (h *ChatbotHandler) Ask(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	role := r.Header.Get("X-User-Role")

	answer, err := h.Service.Ask(services.AskRequest{
		Message: body.Message,
		Role:    role,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"answer": answer})
}

func (h *ChatbotHandler) GetInstructions(w http.ResponseWriter, r *http.Request) {
	list, err := h.Service.Repo.GetAllInstructions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (h *ChatbotHandler) CreateInstruction(w http.ResponseWriter, r *http.Request) {
	var i models.BotInstruction
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if err := h.Service.Repo.CreateInstruction(&i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ChatbotHandler) UpdateInstruction(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var i models.BotInstruction
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	i.ID = uint(id)
	if err := h.Service.Repo.UpdateInstruction(&i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ChatbotHandler) DeleteInstruction(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.Repo.DeleteInstruction(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ChatbotHandler) GetSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := h.Service.Repo.GetSettings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)
}

func (h *ChatbotHandler) UpdateSettings(w http.ResponseWriter, r *http.Request) {
	var s models.BotSettings
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	current, err := h.Service.Repo.GetSettings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.ID = current.ID
	if err := h.Service.Repo.UpdateSettings(&s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
