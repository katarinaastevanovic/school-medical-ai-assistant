package handlers

import (
	"auth-service/dto"
	"auth-service/models"
	"auth-service/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type AuthHandler struct {
	Service services.AuthService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID        uint        `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     string      `json:"email"`
	Role      models.Role `json:"role"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  AuthResponse `json:"user"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.Member

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.Service.Register(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := AuthResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := services.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		Token: token,
		User: AuthResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) GetMemberByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	member, err := h.Service.GetMemberByID(uint(id))
	if err != nil {
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(member)
}

func (h *AuthHandler) GetAllMembers(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllMembers called")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	members, err := h.Service.GetAllMembers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

func (h *AuthHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	members, err := h.Service.GetAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response []dto.MemberResponse
	for _, m := range members {
		response = append(response, dto.MemberResponse{
			ID:        m.ID,
			UMCN:      m.UMCN,
			FirstName: m.FirstName,
			LastName:  m.LastName,
			Email:     m.Email,
			Role:      m.Role,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	members, err := h.Service.GetAllTeachers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response []dto.MemberResponse
	for _, m := range members {
		response = append(response, dto.MemberResponse{
			ID:        m.ID,
			UMCN:      m.UMCN,
			FirstName: m.FirstName,
			LastName:  m.LastName,
			Email:     m.Email,
			Role:      m.Role,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	members, err := h.Service.GetAllDoctors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response []dto.MemberResponse
	for _, m := range members {
		response = append(response, dto.MemberResponse{
			ID:        m.ID,
			UMCN:      m.UMCN,
			FirstName: m.FirstName,
			LastName:  m.LastName,
			Email:     m.Email,
			Role:      m.Role,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) UpdateMember(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.PathValue("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdateMemberRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	member, err := h.Service.UpdateMember(uint(id), &req)
	if err != nil {
		if err.Error() == "member not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(member)
}
