package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/OrigamiKing3612/The-Numbers-App/internal/database"
	"github.com/OrigamiKing3612/The-Numbers-App/internal/routes"
	"github.com/OrigamiKing3612/The-Numbers-App/internal/server"
	"github.com/joho/godotenv"

	_ "modernc.org/sqlite"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	log.Printf("The Numbers App Server started on: http://%s:%s", os.Getenv("HOSTNAME"), os.Getenv("PORT"))
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if debug {
		log.Println("Debug mode is enabled.")
	}

	db, err := sql.Open("sqlite", "./thenumbersapp.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Failed to ping SQLite: %v", err)
	}
	log.Printf("Connected to SQLite!\n")

	queries := database.New(db)

	config := server.NewConfig(debug)

	server := server.NewServer(db, config, queries)

	handler := routes.InitRoutes(server)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: handler,
	}

	go func() {
		log.Println("Ready!")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited gracefully.")
}
