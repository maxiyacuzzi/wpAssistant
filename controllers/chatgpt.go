package controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"wpassistant/config"
	openai "github.com/sashabaranov/go-openai"
)

func HandleChatGPTRequest(c *gin.Context) {
	var request struct {
		Phone   string `json:"phone"`
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido"})
		return
	}

	// Verificar si el usuario existe
	var userID int
	err := config.DB.QueryRow(context.Background(), "SELECT id FROM users WHERE phone=$1", request.Phone).Scan(&userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Conectar con OpenAI
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "Eres un asistente virtual"},
			{Role: "user", Content: request.Message},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en OpenAI"})
		return
	}

	responseText := resp.Choices[0].Message.Content

	// Guardar respuesta en la base de datos
	_, err = config.DB.Exec(context.Background(), "UPDATE messages SET response=$1 WHERE user_id=$2 AND message=$3", responseText, userID, request.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando respuesta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": responseText})
}
