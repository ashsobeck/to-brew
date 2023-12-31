package brews

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"tobrew/types"
)

type Brews struct {
	*types.Server
}

type BrewController interface {
	BrewRoutes() chi.Router
}

func (s *Brews) BrewRoutes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", s.getAllBrews)
	r.Post("/", s.makeNewBrew)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", s.getBrew)
		r.Post("/", s.makeNewBrew)
		r.Put("/", s.updateBrew)
		r.Delete("/", s.deleteBrew)
	})

	return r
}

func (s *Brews) makeNewBrew(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	slog.Info("Request Body: %s", reqBody)
	if err != nil {
		panic(err)
	}

	var brew types.ToBrew
	err = json.Unmarshal(reqBody, &brew)
	if err != nil {
		panic(err)
	}
	if id := chi.URLParam(r, "id"); id != "" {
		brew.Id = id
	} else {
		brew.Id = uuid.NewString()
	}

	tx := s.Db.MustBegin()
	slog.Info("Brew Name: %s", brew.Name)
	slog.Info("Bean: %s", brew.Bean)
	slog.Info("Time: %s", brew.TimeToBrew)
	brewTime, _ := time.Parse(time.RFC3339, brew.TimeToBrew)
	slog.Info("Time: %s", brewTime)
	tx.MustExec(`INSERT INTO tobrews (id, name, bean, roaster, link, brewed, time_of_brew, created)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		brew.Id, brew.Name, brew.Bean, brew.Roaster.String, brew.Link.String, brew.Brewed, brewTime, time.Now().UTC())
	tx.Commit()

	if err = json.NewEncoder(w).Encode(brew); err != nil {
		panic(err)
	}
}

func (s *Brews) deleteBrew(w http.ResponseWriter, r *http.Request) {
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

	tx := s.Db.MustBegin()
	tx.MustExec(`DELETE FROM tobrews WHERE id=?`, id)
	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Brews) getAllBrews(w http.ResponseWriter, r *http.Request) {
	var brews []types.ToBrew
	slog.Info("Getting all brews...")

	err := s.Db.Select(&brews, "SELECT * FROM tobrews ORDER BY time_of_brew DESC")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error(err.Error())
		if _, err = w.Write([]byte(err.Error())); err != nil {
			slog.Error(err.Error())
			panic(err)
		}
	}

	// if len(brews) == 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	if _, err = w.Write([]byte("No brews found.")); err != nil {
	// 		slog.Error(err.Error())
	// 		panic(err)
	// 	}
	// }

	for _, brew := range brews {
		slog.Info(brew.Name)
	}

	if err = json.NewEncoder(w).Encode(brews); err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}

func (s *Brews) getBrew(w http.ResponseWriter, r *http.Request) {
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

	var brew types.ToBrew

	// Get the brew and encode it back to the user
	err := s.Db.Get(&brew, `* FROM tobrews WHERE id=$1`, id)
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

func (s *Brews) updateBrew(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	slog.Info(string(reqBody))
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("Brew not found")); err != nil {
			panic(err)
		}
	}

	var brew types.ToBrew
	err = json.Unmarshal(reqBody, &brew)
	brew.Id = id
	if err != nil {
		panic(err)
	}

	tx := s.Db.MustBegin()
	tx.MustExec(`
        UPDATE tobrews 
        SET name = IFNULL(?, name), bean = IFNULL(?, bean), link = IFNULL(?, link), roaster = IFNULL(?, roaster), brewed = IFNULL(?, brewed)
        WHERE id = ?
    `, brew.Name, brew.Bean, brew.Link, brew.Roaster, brew.Brewed, id)

	if err = tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
	}

	if err = json.NewEncoder(w).Encode(brew); err != nil {
		panic(err)
	}
}
