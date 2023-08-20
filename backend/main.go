package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"tobrew/controllers/brews"
	"tobrew/types/server"
)

func main() {
	enverr := godotenv.Load(".env")
	if enverr != nil {
		slog.Error("ENV gone")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbUser, dbPass, dbHost, dbPort, dbName,
		),
	)
	if err != nil {
		slog.Error(err.Error())
	}
	defer db.Close()

	s := server.Server{Db: db}
	handleRequests(&s)
}

func handleRequests(s *server.Server) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	brewController := brews.Brews{Server: s}

	// r.Post("/new-brew", s.makeNewBrew)
	r.Mount("/tobrews", brewController.BrewRoutes())

	http.ListenAndServe(":3333", r)
}
