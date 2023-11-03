package brews

import (
	"log/slog"
	"net/http"
	"tobrew/types"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	return r
}

func (s *Beans) GetAllBeans(w http.ResponseWriter, r *http.Request) {
	var beans = []types.Bean
	slog.Info("Getting all beans...")

	err := s.Db.Select(&beans, "SELECT * from beans")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//Marshal json here and encode it to send back to client
}
