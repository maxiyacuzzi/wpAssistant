package config

import (
	"context"
	"os"
	"github.com/sashabaranov/go-openai"
)

// Crear una instancia global del cliente de OpenAI
var OpenAIClient *openai.Client

func ConnectOpenAI() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	OpenAIClient = openai.NewClient(apiKey)
}
