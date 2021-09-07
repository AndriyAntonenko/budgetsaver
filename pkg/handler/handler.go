package handler

import (
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

	r.Post("/auth/sign-up", h.createUser)

	return r
}
