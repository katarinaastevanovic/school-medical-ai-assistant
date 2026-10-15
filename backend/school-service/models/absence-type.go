package models

type AbsenceType string

const (
	Pending   AbsenceType = "PENDING"
	Excused   AbsenceType = "EXCUSED"
	Unexcused AbsenceType = "UNEXCUSED"
)
