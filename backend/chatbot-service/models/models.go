package models

import "gorm.io/gorm"

// BotInstruction su "gradivni blokovi" system prompta koje admin uređuje
type BotInstruction struct {
	gorm.Model
	Content string `json:"content"`
	Role    string `json:"role"` // "", "STUDENT", "TEACHER", "DOCTOR", "ADMIN" - "" = za sve
	Order   int    `json:"order"`
	Enabled bool   `json:"enabled" gorm:"default:true"`
}

// BotSettings je jedan red - globalna podešavanja generisanja odgovora
type BotSettings struct {
	gorm.Model
	Temperature float32 `json:"temperature" gorm:"default:0.7"`
	MaxTokens   int     `json:"maxTokens" gorm:"default:512"`
	GroqModel   string  `json:"groqModel" gorm:"default:'llama-3.3-70b-versatile'"`
}
