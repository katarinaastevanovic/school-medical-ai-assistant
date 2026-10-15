package dto

import "auth-service/models"

type MemberResponse struct {
	ID        uint        `json:"id"`
	UMCN      string      `json:"umcn"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     string      `json:"email"`
	Role      models.Role `json:"role"`
}

type UpdateMemberRequest struct {
	UMCN      string      `json:"umcn"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Role      models.Role `json:"role"`
}
