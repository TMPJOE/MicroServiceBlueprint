package main

import (
	"net/http"
	"os"

	"hotel.com/app/internal/database"
	"hotel.com/app/internal/handler"
	"hotel.com/app/internal/logging"
	"hotel.com/app/internal/repo"
	"hotel.com/app/internal/service"
)

func main() {
	//create logger
	l := logging.New()
	l.Info("App initiated")

	//db connection
	db, err := database.NewConn()
	if err != nil {
		l.Error("Conection to database failed", "err", err)
		os.Exit(-1)
	}
	l.Info("Database connection successful")

	defer db.Close()

	//repo creation
	r := repo.New(db)

	//service creation
	svc := service.New(l, r)

	//hanlder creation
	h := handler.New(svc, l)

	//server creation
	svr := h.NewServerMux()

	//serve microservice
	http.ListenAndServe("localhost:8000", svr)
}
