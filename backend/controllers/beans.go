package brews

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"tobrew/types"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type Beans struct {
	*types.Server
}

func (s *Beans) BeanRoutes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", s.GetAllBeans)
	r.Post("/", s.CreateNewBean)

	r.Route("/{id}", func(r chi.Router) {
		r.Delete("/", s.DeleteBean)
	})

	return r
}

func (s *Beans) GetAllBeans(w http.ResponseWriter, r *http.Request) {
	var beans []types.Bean
	slog.Info("Getting all beans...")

	err := s.Db.Select(&beans, "SELECT * from beans")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Marshal json here and encode it to send back to client
	if err = json.NewEncoder(w).Encode(beans); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Beans) CreateNewBean(w http.ResponseWriter, r *http.Request) {
	var bean types.Bean
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(reqBody, &bean); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	slog.Info("Creating new bean")
	bean.Id = uuid.NewString()

	tx := s.Db.MustBegin()
	insertBean := `INSERT INTO beans (Id, Name, Roaster, Country, Varietal, Process, Altitude, Notes, Weight)  
				   VALUES (?,?,?,?,?,?,?,?,?)`
	tx.MustExec(insertBean, bean.Id, bean.Name, bean.Roaster, bean.Country, bean.Varietal, bean.Process, bean.Altitude, bean.Notes, bean.Weight)

	if err = tx.Commit(); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(bean); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Beans) DeleteBean(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tx := s.Db.MustBegin()
	deleteBean := `DELETE FROM beans WHERE id = ?`
	tx.MustExec(deleteBean, id)
	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
