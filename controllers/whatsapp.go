package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"wpassistant/config"
)

func HandleWhatsAppMessage(c *gin.Context) {
	var message struct {
		From    string `json:"from"`
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido"})
		return
	}

	// Verificar si el usuario existe
	var userID int
	err := config.DB.QueryRow(context.Background(), "SELECT id FROM users WHERE phone=$1", message.From).Scan(&userID)
	if err != nil {
		// Si el usuario no existe, crearlo
		err = config.DB.QueryRow(context.Background(), "INSERT INTO users (phone) VALUES ($1) RETURNING id", message.From).Scan(&userID)
		if err != nil {
			log.Println("Error creando usuario:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando usuario"})
			return
		}
	}

	// Guardar mensaje
	_, err = config.DB.Exec(context.Background(), "INSERT INTO messages (user_id, message) VALUES ($1, $2)", userID, message.Message)
	if err != nil {
		log.Println("Error guardando mensaje:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando mensaje"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Mensaje recibido"})
}
