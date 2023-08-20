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

func (s *Server) makeNewBrew(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var brew ToBrew
	err = json.Unmarshal(reqBody, &brew)
	if err != nil {
		panic(err)
	}

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

	if err = json.NewEncoder(w).Encode(brew); err != nil {
		panic(err)
	}
}

func (s *Server) getBrew(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id != "" {
		slog.Info("Getting Brew with Id: %s", id)

		var brew ToBrew

		// Get the brew and encode it back to the user
		err := s.db.Get(&brew, `SELECT * FROM tobrews WHERE id=$1`, id)
		if err != nil {
			w.WriteHeader(500)
			if _, err = w.Write([]byte(err.Error())); err != nil {
				panic(err)
			}
		}

		if err = json.NewEncoder(w).Encode(brew); err != nil {
			panic(err)
		}

	}
	// brew isn't found
	w.WriteHeader(404)
	if _, err := w.Write([]byte("Brew not found")); err != nil {
		panic(err)
	}
}

func (s *Server) getAllBrews(w http.ResponseWriter, r *http.Request) {
	var brews []ToBrew

	err := s.db.Select(&brews, "SELECT * FROM tobrews ORDER BY time_of_brew DESC")
	if err != nil {
		w.WriteHeader(500)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
	}

	if len(brews) == 0 {
		w.WriteHeader(404)
		if _, err = w.Write([]byte("No brews found.")); err != nil {
			panic(err)
		}
	}

	if err = json.NewEncoder(w).Encode(brews); err != nil {
		panic(err)
	}
}

func (s *Server) brewRoutes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", s.getAllBrews)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", s.getBrew)
		r.Post("/", s.makeNewBrew)
	})

	return r
}

func (s *Server) handleRequests() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	r.Post("/new-brew", s.makeNewBrew)
	r.Mount("/tobrews", s.brewRoutes())

	http.ListenAndServe(":3333", r)
}
