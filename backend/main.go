package main

import (
	"context"
	"exchangeapp/backend/config"
	"exchangeapp/backend/router"
	"exchangeapp/backend/validation"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin/binding"
)

func main() {
	config.InitConfig()
	binding.Validator = validation.NewValidator()
	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

	// r.Run("0.0.0.0" + port)
}
