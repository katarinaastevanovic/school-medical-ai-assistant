package dto

type StudentSummaryResponse struct {
	AverageGrade float64 `json:"averageGrade"`
	Excused      int64   `json:"excused"`
	Unexcused    int64   `json:"unexcused"`
	Pending      int64   `json:"pending"`
}
