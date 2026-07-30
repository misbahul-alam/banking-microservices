package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/misbahul-alam/banking-microservices/services/auth/internal/config"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/database"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	logr, err := logger.New()
	if err != nil {
		log.Fatal(err)
	}

	defer logr.Sync()

	if err := database.RunMigrations(cfg); err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logr.Info("Database connection established")

	logr.Info("Auth service starting...",
		zap.String("app_name", cfg.AppName),
		zap.String("app_env", cfg.AppEnv),
		zap.String("grpc_port", cfg.GRPCPort))

	log.Println("Starting API Gateway...")
	log.Println("API Gateway running on port 8080. Routing requests...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down API Gateway...")

}
