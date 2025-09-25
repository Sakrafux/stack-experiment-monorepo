package rest

import (
	"net/http"
	"time"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/config"
	mw "github.com/Sakrafux/stack-experiment-monorepo/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

type Api struct {
	config *config.Config
	db     *sqlx.DB
}

func NewApi(config *config.Config, db *sqlx.DB) *Api {
	return &Api{config, db}
}

func (api *Api) CreateRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(mw.Logging)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		userApi := NewUserApi(api)
		r.Mount("/users", userApi.CreateUsersRouter())
		r.Mount("/user", userApi.CreateUserRouter())
	})

	return r
}
