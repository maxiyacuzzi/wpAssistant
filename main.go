package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"wpassistant/config" 
	"wpassistant/routes"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Conectar a la base de datos
	config.ConnectDB()

	// Inicializar servidor
	r := gin.Default()

	// Definir rutas
	routes.SetupRoutes(r)

	// Obtener el puerto de las variables de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Iniciar el servidor
	log.Println("Servidor corriendo en el puerto " + port)
	r.Run(":" + port)
}
