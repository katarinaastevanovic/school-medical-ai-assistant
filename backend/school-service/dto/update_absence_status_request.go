package dto

import "school-service/models"

type UpdateAbsenceStatusRequest struct {
	Status models.AbsenceType `json:"status"`
}