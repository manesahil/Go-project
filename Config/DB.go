package Config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// Get connection string from env
	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		log.Fatal("DB_CONN_STR is not set")
	}

	// Connect to DB
	DB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	// Ping DB
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatal("Database not reachable:", err)
	}

	log.Println("Database connected successfully!")

}
