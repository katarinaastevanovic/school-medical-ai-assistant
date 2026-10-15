package services

import (
	"chatbot-service/groq"
	"chatbot-service/repository"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type ChatbotService struct {
	Repo repository.ChatbotRepository
}

type AskRequest struct {
	Message string
	Role    string
}

func (s *ChatbotService) Ask(req AskRequest) (string, error) {
	instructions, err := s.Repo.GetActiveInstructions(req.Role)
	if err != nil {
		return "", err
	}

	settings, err := s.Repo.GetSettings()
	if err != nil {
		return "", err
	}

	contextData := s.gatherContext()

	var sb strings.Builder
	for _, instr := range instructions {
		sb.WriteString(instr.Content)
		sb.WriteString("\n")
	}
	sb.WriteString("\nPodaci koje mozes da koristis za odgovor (nemoj izmisljati podatke van ovoga):\n")
	sb.WriteString(contextData)

	return groq.Ask(sb.String(), req.Message, settings.GroqModel, settings.Temperature, settings.MaxTokens)
}

func (s *ChatbotService) gatherContext() string {
	var parts []string

	schoolURL := os.Getenv("SCHOOL_SERVICE_URL")
	if resp, err := http.Get(schoolURL + "/internal/chatbot-context"); err == nil {
		defer resp.Body.Close()
		var data map[string]interface{}
		if json.NewDecoder(resp.Body).Decode(&data) == nil {
			b, _ := json.Marshal(data)
			parts = append(parts, "Skolski podaci: "+string(b))
		}
	}

	medicalURL := os.Getenv("MEDICAL_SERVICE_URL")
	if resp, err := http.Get(medicalURL + "/internal/chatbot-context"); err == nil {
		defer resp.Body.Close()
		var data map[string]interface{}
		if json.NewDecoder(resp.Body).Decode(&data) == nil {
			b, _ := json.Marshal(data)
			parts = append(parts, "Medicinski podaci: "+string(b))
		}
	}

	return strings.Join(parts, "\n")
}
