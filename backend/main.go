package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type ToBrew struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Bean       string `json:"bean"`
	TimeToBrew string `json:"timeToBrew"`
}

type Server struct {
	db *sqlx.DB
}

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

	s := Server{db: db}
	s.handleRequests()
}

func (s *Server) makeNewToBrew(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var brew ToBrew
	json.Unmarshal(reqBody, &brew)

	if brew.Id == "" {
		brew.Id = uuid.NewString()
	}

	tx := s.db.MustBegin()
	slog.Info("Brew Name: %s", brew.Name)
	slog.Info("Bean: %s", brew.Bean)
	brewTime, _ := time.Parse(time.RFC3339, brew.TimeToBrew)
	tx.MustExec(`INSERT INTO tobrews (id, name, bean, time_of_brew, created)
        VALUES (?, ?, ?, ?, ?)`,
		brew.Id, brew.Name, brew.Bean, brewTime, time.Now().UTC())
	tx.Commit()

	json.NewEncoder(w).Encode(brew)
}

func (s *Server) handleRequests() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	r.Post("/new-brew", s.makeNewToBrew)

	http.ListenAndServe(":3333", r)
}
