package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Moosa-Razaa/Authentication/internal/database"
)

type App struct {
	server *http.Server
}

func NewApplication() *App {
	return &App{
		server: &http.Server{
			Addr: ":8080",
		},
	}
}

func (application *App) Start() error {
	if err := database.InitializeDatabase(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	defer database.CloseDatabase()

	router := SetupRoutes()
	application.server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := application.server.ListenAndServe(); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	fmt.Println("Server started on port 8080")
	application.waitForShutdown()
	return nil
}

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service is healthy"))
	})

	return mux
}

func (application *App) waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")
	if err := application.server.Shutdown(context.TODO()); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	} else {
		fmt.Println("Server shut down gracefully")
	}

	fmt.Println("Closing database connection...")
	if err := database.CloseDatabase(); err != nil {
		fmt.Printf("Error closing database connection: %v\n", err)
	} else {
		fmt.Println("Database connection closed")
	}

	fmt.Println("Application terminated")
}
