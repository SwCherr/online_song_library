package main

import (
	"context"
	"online_music/base"
	"online_music/pkg/handler"
	"online_music/pkg/repository"
	"online_music/pkg/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title           Online song library API
// @version         1.0
// @description     API server for online song library
// @host      localhost:8000
// @BasePath  /api

func main() {
	settingLogs()
	initVariableENV()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialization DB: %s", err.Error())
	}

	logrus.Info("DB is initializate")

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	srv := new(base.Server)

	go func() {
		if err := srv.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error ocurred while running HTTP server: %s", err.Error())
		}
		logrus.Info("Server is run")
	}()

	// down without dropping any requests
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("online_music is shutting down")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error shutting down database: %s", err.Error())
	}
}

func settingLogs() {
	err := os.Mkdir("logs", 0777)
	if err != nil {
		logrus.Info("error created dir for logs: ", err.Error())
	}

	file, err := os.OpenFile("logs/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func initVariableENV() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variable: %s", err.Error())
	}
	logrus.Info("loaded ENV variable")
}
