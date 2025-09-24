package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/db"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/user"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type UserApi struct {
	service *user.Service
}

func NewUserApi(sqlDb *sqlx.DB) *UserApi {
	return &UserApi{user.NewService(db.NewUserRepository(sqlDb))}
}

func (api *UserApi) CreateUsersRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /", api.RegisterUser)

	return router
}

func (api *UserApi) CreateUserRouter() *http.ServeMux {
	router := http.NewServeMux()

	return router
}

func (api *UserApi) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.HandleHttpError(w, errors.NewBadRequestError(err.Error()))
		return
	}

	u, err := api.service.RegisterUser(r.Context(), toUser(req.User))
	if err != nil {
		errors.HandleHttpError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(UserResponse{fromUser(u)})
	if err != nil {
		errors.HandleHttpError(w, err)
		return
	}
}
