package user

import (
	"encoding/json"
	"net/http"

	"github.com/KornutaKM/news_aggregator_with_tg/internal/auth"
	"github.com/KornutaKM/news_aggregator_with_tg/pkg/response"
	"github.com/KornutaKM/news_aggregator_with_tg/pkg/validator"
)

type Handler struct {
	service    *Service
	jwtService *auth.JWT
}

func NewHandler(service *Service, router *http.ServeMux, jwtService *auth.JWT) *Handler {
	h := &Handler{service: service, jwtService: jwtService}

	router.HandleFunc("/auth/register", h.Register())
	router.HandleFunc("/auth/login", h.Login())

	return h
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterUserRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.RespondWithError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if err := validator.ValidateStruct(req); err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := h.service.Register(req.Email, req.Password)
		if err != nil {
			if err == ErrEmailAlreayExists {
				response.RespondWithError(w, http.StatusConflict, "email already exists")
				return
			}
			response.RespondWithError(w, http.StatusInternalServerError, "failed to register user")
			return
		}

		response.RespondWithJSON(w, http.StatusCreated, user)
	}
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginUserRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.RespondWithError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if err := validator.ValidateStruct(req); err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := h.service.Login(req.Email, req.Password)
		if err != nil {
			response.RespondWithError(w, http.StatusUnauthorized, "invalid email or password")
			return
		}

		token, err := h.jwtService.GenerateToken(user.ID, "user")
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, "failed to generate token")
			return
		}

		response.RespondWithJSON(w, http.StatusOK, map[string]any{
			"token": token,
			"user":  user,
		})
	}
}
