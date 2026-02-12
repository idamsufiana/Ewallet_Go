package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"wallet-service/internal/database"
	"wallet-service/internal/handler"
	"wallet-service/internal/repository"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load database
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer db.Close()

	// Dependency Injection
	userRepo := repository.NewPostgresUserRepository(db)
	walletService := service.NewWalletService(db, userRepo)
	walletHandler := handler.NewWalletHandler(walletService)

	// Router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.POST("/withdraw", walletHandler.Withdraw)
	router.GET("/balance/:user_id", walletHandler.GetBalance)

	// Start server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("server error:", err)
		}
	}()

	log.Println("Server running on port 8080")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
