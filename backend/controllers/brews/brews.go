package brew

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	. "tobrew/types/server"
	. "tobrew/types/tobrew"
)

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

	tx := s.Db.MustBegin()
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
	// chi.URLParam gets the variables from the route NOT the query param
	// ie: /brew/1234 where 1234 is id
	id := chi.URLParam(r, "id")

	if id == "" {
		// brew isn't found
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("Brew not found")); err != nil {
			panic(err)
		}
	}

	slog.Info("Getting Brew with Id: %s", id)

	tx := s.db.MustBegin()
	tx.MustExec(`DELETE FROM tobrews WHERE id=?`, id)
	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getAllBrews(w http.ResponseWriter, r *http.Request) {
	var brews []ToBrew

	err := s.db.Select(&brews, "SELECT * FROM tobrews ORDER BY time_of_brew DESC")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
	}

	if len(brews) == 0 {
		w.WriteHeader(http.StatusNotFound)
		if _, err = w.Write([]byte("No brews found.")); err != nil {
			panic(err)
		}
	}

	if err = json.NewEncoder(w).Encode(brews); err != nil {
		panic(err)
	}
}

func (s *Server) deleteBrew(w http.ResponseWriter, r *http.Request) {
	// chi.URLParam gets the variables from the route NOT the query param
	// ie: /brew/1234 where 1234 is id
	id := chi.URLParam(r, "id")

	if id == "" {
		// brew isn't found
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("Brew not found")); err != nil {
			panic(err)
		}
	}

	slog.Info("Deleting Brew with Id: %s", id)

	var brew ToBrew

	// Get the brew and encode it back to the user
	err := s.db.Get(&brew, `* FROM tobrews WHERE id=$1`, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
	}

	if err = json.NewEncoder(w).Encode(brew); err != nil {
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
		r.Delete("/", s.deleteBrew)
	})

	return r
}
