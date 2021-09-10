package handler

import (
	"errors"
	"net/http"
	"strings"

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
