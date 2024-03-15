package brews

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"tobrew/types"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type Brews struct {
	*types.Server
}

type BrewController interface {
	BrewRoutes() chi.Router
}

type brewIdWeight struct {
	Id     string
	Weight float32
}

type BrewResponse struct {
	Brew      types.ToBrew `json:"brew"`
	NewWeight float32      `json:"newWeight"`
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
	r.Route("/complete/{id}", func(r chi.Router) {
		r.Put("/", s.markBrewed)
	})

	return r
}

func (s *Brews) makeNewBrew(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	slog.Info("Request Body: %s", reqBody)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	var brew types.ToBrew
	err = json.Unmarshal(reqBody, &brew)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	if id := chi.URLParam(r, "id"); id != "" {
		brew.Id = id
	} else {
		brew.Id = uuid.NewString()
	}

	tx := s.Db.MustBegin()
	slog.Info("Brew Name:", brew.Name)
	slog.Info("Bean:", brew.Bean)
	slog.Info("Time:", brew.TimeToBrew)
	brewTime, _ := time.Parse(time.RFC3339, brew.TimeToBrew)
	slog.Info("Time:", brewTime)
	tx.MustExec(`INSERT INTO tobrews (id, name, bean, roaster, link, brewed, time_of_brew, created, BeanWeight)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		brew.Id, brew.Name, brew.Bean, brew.Roaster.String, brew.Link.String, brew.Brewed, brewTime, time.Now().UTC(), brew.BeanWeight)
	tx.Commit()

	if err = json.NewEncoder(w).Encode(brew); err != nil {
		slog.Error(err.Error())
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
	slog.Info(brew.Bean)
	tx.MustExec(`
        UPDATE tobrews 
        SET name = IF(? != "", name, ?), bean = IF(? != "", ?, bean), link = IF(? != "", ?, link), roaster = IF(? != "", ?, roaster), brewed = IFNULL(?, brewed)
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

func (s *Brews) markBrewed(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	slog.Info(id)
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte("Brew not found")); err != nil {
			panic(err)
		}
	}

	slog.Info("Marking as brewed...")
	tx := s.Db.MustBegin()
	tx.MustExec(`UPDATE tobrews SET brewed = true WHERE id = ?`, id)
	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error(err.Error())
		if _, err = w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
		return
	}

	slog.Info("Getting brew...")

	var brew types.ToBrew
	if err := s.Db.Get(&brew, `SELECT * FROM tobrews WHERE id = ? LIMIT 1`, id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error(err.Error())
		if _, err = w.Write([]byte(err.Error())); err != nil {
			panic(err)
		}
		return
	}
	slog.Info("brew:", brew)

	beanService := Beans{Server: s.Server}
	var res BrewResponse
	res.Brew = brew

	if newWeight, err := beanService.Brew(brew.Bean, brew.BeanWeight); err == nil {
		res.NewWeight = newWeight
		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			slog.Error(err.Error())

			if _, err = w.Write([]byte(err.Error())); err != nil {
				panic(err)
			}
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if _, err = w.Write([]byte("Bean not found")); err != nil {
			panic(err)
		}

	}
}
