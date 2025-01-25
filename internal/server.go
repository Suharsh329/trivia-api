package internal

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	config "trivia/internal/config"
	routes "trivia/internal/routes"
)

func Run() {
	// Load the env file
	config.LoadEnv()

	db := config.ConnectDatabase()

	defer db.Close()

	mux := http.NewServeMux()

	cors := config.Cors()

	routes.RegisterRoutes(mux, db)

	port := ":" + config.Config("PORT")

	server := &http.Server{
		Addr:    port,
		Handler: cors.Handler(mux),
	}

	go func() {
		if config.Config("APP_ENV") == "local" {
			log.Printf("Server started: http://localhost%s\n", server.Addr)
		}
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
