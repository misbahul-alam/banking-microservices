package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/misbahul-alam/banking-microservices/services/auth/internal/config"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/database"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/database/sqlc"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/dto"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/logger"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/password"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/repository"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/service"
	"github.com/misbahul-alam/banking-microservices/services/auth/internal/token"
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

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := sqlc.New(db)
	repo := repository.NewUserRepository(queries)
	pass := password.New(10)
	tkn, err := token.NewJWTMaker("SEC")
	srv := service.NewAuthService(repo, pass, tkn, time.Minute*15, time.Hour*24*30)

	res, err := srv.Login(context.Background(), dto.LoginRequest{
		Email:    "admin@email.com",
		Password: "12345678",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

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
