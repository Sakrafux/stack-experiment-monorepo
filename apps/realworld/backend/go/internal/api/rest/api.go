package rest

import (
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/config"
	"github.com/jmoiron/sqlx"
)

type Api struct {
	config config.Config
	db     *sqlx.DB
}

func NewApi(config config.Config, db *sqlx.DB) *Api {
	return &Api{config, db}
}

func (api *Api) CreateRouter() *http.ServeMux {
	router := http.NewServeMux()

	userApi := NewUserApi(api.db)
	router.Handle("/users", userApi.CreateUsersRouter())
	router.Handle("/user", userApi.CreateUserRouter())

	routerWrapper := http.NewServeMux()
	routerWrapper.Handle("/api/", http.StripPrefix("/api", router))
	return routerWrapper
}
