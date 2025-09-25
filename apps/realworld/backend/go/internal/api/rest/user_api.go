package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/db"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/user"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/errors"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/middleware"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/security"
	"github.com/go-chi/chi/v5"
)

const REFRESH_COOKIE_NAME = "real_world-refresh_token"

type UserApi struct {
	api     *Api
	repo    *db.UserRepository
	service *user.Service
}

func NewUserApi(api *Api) *UserApi {
	repo := db.NewUserRepository(api.db)
	service := user.NewService(api.config, repo)
	return &UserApi{api, repo, service}
}

func (api *UserApi) CreateUsersRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/", api.RegisterUser)
	r.Post("/login", api.Login)

	return r
}

func (api *UserApi) CreateUserRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", api.GetUser)
	r.Put("/", api.PutUser)
	r.Get("/token", api.RefreshToken)

	return r
}

func (api *UserApi) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.HandleHttpError(w, r, errors.NewBadRequestError(err.Error()))
		return
	}

	u, err := api.service.RegisterUser(r.Context(), fromNewUser(req.User))
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
	userResponse := UserResponse{toUser(u)}

	token, err := api.createAccessToken(u.Id)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
	userResponse.User.Token = token

	err = api.createRefreshTokenCookie(w, u.Id)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(userResponse)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
}

func (api *UserApi) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.HandleHttpError(w, r, errors.NewBadRequestError(err.Error()))
		return
	}

	u, err := api.service.LoginUser(r.Context(), fromLoginUser(req.User))
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
	userResponse := UserResponse{toUser(u)}

	token, err := api.createAccessToken(u.Id)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
	userResponse.User.Token = token

	err = api.createRefreshTokenCookie(w, u.Id)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(userResponse)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
}

func (api *UserApi) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	if !ok {
		errors.HandleHttpError(w, r, errors.NewUnauthorizedError("user is not authenticated"))
		return
	}

	u, err := api.service.FindUserById(ctx, userId)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
	userResponse := UserResponse{toUser(u)}

	token, err := api.createAccessToken(u.Id)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
	userResponse.User.Token = token

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(userResponse)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
}

func (api *UserApi) PutUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	if !ok {
		errors.HandleHttpError(w, r, errors.NewUnauthorizedError("user is not authenticated"))
		return
	}

	var req UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.HandleHttpError(w, r, errors.NewBadRequestError(err.Error()))
		return
	}

	serviceUser := fromUpdateUser(req.User)
	serviceUser.Id = userId

	u, err := api.service.UpdateUser(ctx, serviceUser)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
	userResponse := UserResponse{toUser(u)}

	token, err := api.createAccessToken(u.Id)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
	userResponse.User.Token = token

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(userResponse)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
}

func (api *UserApi) RefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(REFRESH_COOKIE_NAME)
	if err != nil {
		errors.HandleHttpError(w, r, errors.NewBadRequestError(err.Error()))
		return
	}
	refreshToken := cookie.Value

	validatedToken, err := security.ValidateRefreshToken(refreshToken, api.api.config.JWT.RefreshSecret)
	if err != nil {
		errors.HandleHttpError(w, r, errors.NewUnauthorizedError(err.Error()))
		return
	}
	data, err := security.ExtractTokenData(validatedToken)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	token, err := api.createAccessToken(data.Id)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	err = api.createRefreshTokenCookie(w, data.Id)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	_, err = w.Write([]byte(token))
	if err != nil {
		errors.HandleHttpError(w, r, err)
	}
}

func (api *UserApi) createAccessToken(userId int64) (string, error) {
	accessToken, err := security.CreateAccessToken(&security.TokenData{Id: userId, Role: "user"}, api.api.config.JWT.AccessSecret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (api *UserApi) createRefreshTokenCookie(w http.ResponseWriter, userId int64) error {
	refreshToken, err := security.CreateRefreshToken(&security.TokenData{Id: userId, Role: "user"}, api.api.config.JWT.RefreshSecret)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     REFRESH_COOKIE_NAME,
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   false,
		Path:     "/api/user/token",
		MaxAge:   30 * 24 * 60 * 60, // 7 days
		SameSite: http.SameSiteStrictMode,
	})

	return nil
}
