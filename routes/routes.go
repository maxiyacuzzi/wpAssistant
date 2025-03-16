package routes

import (
	"github.com/gin-gonic/gin"
	"wpassistant/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/whatsapp", controllers.HandleWhatsAppMessage)
	r.POST("/chatgpt", controllers.HandleChatGPTRequest)
}
