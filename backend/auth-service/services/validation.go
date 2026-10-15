package services

import (
	"net/mail"
	"strconv"
	"time"
	"unicode"
)

func ValidateJMBG(jmbg string) bool {

	// mora imati 13 cifara
	if len(jmbg) != 13 {
		return false
	}

	// sve mora biti broj
	for _, char := range jmbg {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	// datum
	day, _ := strconv.Atoi(jmbg[0:2])
	month, _ := strconv.Atoi(jmbg[2:4])
	yearPart, _ := strconv.Atoi(jmbg[4:7])

	var year int

	if yearPart >= 900 {
		year = 1000 + yearPart // (retki slučajevi)
	} else {
		year = 2000 + yearPart // moderni JMBG
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	if date.Day() != day ||
		int(date.Month()) != month ||
		date.Year() != year {
		return false
	}

	if date.After(time.Now()) {
		return false
	}

	if year < 1900 {
		return false
	}

	// region
	_ = jmbg[7:9]

	return true
}

func GetGenderFromJMBG(jmbg string) string {
	bbb, _ := strconv.Atoi(jmbg[9:12])

	if len(jmbg) != 13 {
		return ""
	}

	if bbb < 500 {
		return "MALE"
	}

	return "FEMALE"
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
