package main

import (
	"fmt"
	"net/http"
	"os"

	"hotel.com/app/internal/database"
	"hotel.com/app/internal/handler"
	"hotel.com/app/internal/helper"
	"hotel.com/app/internal/logging"
	"hotel.com/app/internal/repo"
	"hotel.com/app/internal/service"
)

type Config struct {
	Port     string
	DbUrl    string
	LogLevel string
}

func main() {
	config := loadConfig()

	//create logger
	l := logging.New()
	l.Info("App initiated")

	//db connection
	db, err := database.NewConn(config.DbUrl)
	if err != nil {
		l.Error("Conection to database failed", "err", err)
		os.Exit(-1)
	}
	l.Info("Database connection successful")

	defer db.Close()

	//repo creation
	r := repo.NewDatabaseRepo(db)

	//service creation
	svc := service.New(l, r)

	//hanlder creation
	h := handler.New(svc, l)

	//server creation
	mux := h.NewServerMux()
	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", config.Port),
		Handler: mux,
	}
	go srv.ListenAndServe()
	// REMEMBER TO USE TLS ON "PRODUCTION"
}

func loadConfig() *Config {
	return &Config{
		Port:     helper.Getenv("APP_PORT", "8080"),
		LogLevel: helper.Getenv("LOG_LEVEL", "0"),
		DbUrl:    os.Getenv("DB_URL"),
	}
}
