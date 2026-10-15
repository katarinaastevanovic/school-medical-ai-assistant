package dto

type AbsenceStatsResponse struct {
	Pending   int64 `json:"pending"`
	Excused   int64 `json:"excused"`
	Unexcused int64 `json:"unexcused"`
}
