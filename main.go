package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ShadrackAdwera/go-payments/api"
	"github.com/ShadrackAdwera/go-payments/authenticator"
	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// vocek57061@oniecan.com
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}
	url := os.Getenv("PG_URL")
	conn, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatalf("Failed to initialize the database %v", err)
	}

	store := db.NewStore(conn)

	srv := api.NewServer(store, auth)

	err = srv.StartServer("0.0.0.0:3000")

	if err != nil {
		panic(err)
	}
}
