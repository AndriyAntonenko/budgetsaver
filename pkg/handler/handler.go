package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	service "github.com/AndriyAntonenko/budgetSaver/pkg/services"
	"github.com/AndriyAntonenko/goRouter"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{srv}
}

func (h *Handler) InitRoutes() *goRouter.Router {
	r := goRouter.NewRouter()

	r.Post("/api/auth/sign-up", h.createUser)
	r.Post("/api/auth/login", h.login)
	r.Get("/api/auth/me", h.me)

	r.Post("/api/group", h.createGroup)

	return r
}

func extractToken(r *http.Request) (string, error) {
	authHeaderValue := r.Header.Get("Authorization")
	if len(authHeaderValue) == 0 {
		return "", errors.New("no auth header")
	}

	splitHeader := strings.Split(authHeaderValue, " ")
	if len(splitHeader) != 2 {
		return "", errors.New("invalid auth header")
	}

	return splitHeader[1], nil
}

func (h *Handler) getUserId(w http.ResponseWriter, r *http.Request) (string, error) {
	accessToken, err := extractToken(r)
	if err != nil {
		logger.UseBasicLogger().Error("Unauthorized error", err, "func me()")
		w.WriteHeader(http.StatusUnauthorized)
		return "", err
	}

	userId, err := h.service.Authorization.ParseAccessToken(accessToken)
	if err != nil {
		logger.UseBasicLogger().Error("Unauthorized error", err, "func me()")
		w.WriteHeader(http.StatusUnauthorized)
		return "", err
	}

	return userId, err
}

func (h *Handler) sendJSON(w http.ResponseWriter, payload interface{}) {
	responseBody, err := json.Marshal(payload)

	if err != nil {
		logger.UseBasicLogger().Error("Internal server error", err, "func createUser()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (h *Handler) parseJSONBody(w http.ResponseWriter, r *http.Request, pt interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pt)
	if err != nil {
		logger.UseBasicLogger().Error("Bad request", err, "func createGroup()")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
