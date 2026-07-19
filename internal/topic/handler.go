package topic

import (
	"encoding/json"
	"net/http"

	"github.com/KornutaKM/news_aggregator_with_tg/internal/middleware"
	"github.com/KornutaKM/news_aggregator_with_tg/pkg/response"
	"github.com/KornutaKM/news_aggregator_with_tg/pkg/validator"
)

type Handler struct {
	service        *Service
	authMiddleware middleware.Middleware
}

func NewHandler(service *Service, router *http.ServeMux, authMiddleware middleware.Middleware) *Handler {
	h := &Handler{service: service}

	router.Handle("POST /topics", authMiddleware(h.Create()))
	router.HandleFunc("GET /topics", h.GetAll())

	return h
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateTopicRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.RespondWithError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if err := validator.ValidateStruct(req); err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		role, ok := middleware.GetUserRoleFromContext(r.Context())
		if !ok || role != "admin" {
			response.RespondWithError(w, http.StatusUnauthorized, "unathorized")
			return
		}

		topic, err := h.service.Create(req.Name)
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, "failed to create topic")
			return
		}

		response.RespondWithJSON(w, http.StatusCreated, topic)
	}
}

func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		topics, err := h.service.GetAll()
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, "failed to get topics")
			return
		}

		response.RespondWithJSON(w, http.StatusOK, topics)
	}
}
