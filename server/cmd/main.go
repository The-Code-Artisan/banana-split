package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	"github.com/the-code-artisan/banana-split/server/internal/repository"
)

func main() {
	// DB Init
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v\n", err)
	}
	defer dbPool.Close()

	repo := repository.New(dbPool)

	users, _ := repo.FindAllUsers(context.Background())
	log.Printf("Users: %v", users)

	router := http.NewServeMux()
	router.HandleFunc("GET /ping", ping)

	// Server Init
	server := http.Server{
		Addr:    ":8001",
		Handler: router,
	}
	log.Println("Starting server on port :8001")
	err = server.ListenAndServe()
	if err != nil {
		log.Panicf("Unable to start server. Err: %v", err)
	}
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Pong!")
}
