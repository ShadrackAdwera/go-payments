package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ShadrackAdwera/go-payments/api"
	"github.com/ShadrackAdwera/go-payments/authenticator"
	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	"github.com/ShadrackAdwera/go-payments/worker"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	zerolog "github.com/rs/zerolog/log"

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

	redisAddress := os.Getenv("REDIS_ADDRESS")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	redisOpts := asynq.RedisClientOpt{
		Addr: redisAddress,
	}

	store := db.NewStore(conn)
	distro := worker.NewTaskDistributor(redisOpts)
	srv := api.NewServer(store, auth, distro)

	go startTaskProcessor(redisOpts, store)

	err = srv.StartServer(serverAddress)

	if err != nil {
		panic(err)
	}
}

func startTaskProcessor(opts asynq.RedisClientOpt, store db.TxStore) {
	processor := worker.NewTaskServer(opts, store)

	err := processor.Start()

	if err != nil {
		zerolog.Err(err).Str("error", "error starting the redis task processor")
		return
	}
	zerolog.Info().Str("start", "redis task processor started")
}

// func seedDbWithPermissionData() {
// 	permissions := utils.GetPermissionData()

// }
