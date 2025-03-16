package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	dbURL := os.Getenv("POSTGRES_URI")
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Error conectando a PostgreSQL:", err)
	}

	fmt.Println("✅ Conectado a PostgreSQL")
	DB = conn
}
