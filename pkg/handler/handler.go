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

	r.EnableCors(&goRouter.RouterCors{
		Origins: "*",
		Methods: "GET, POST, DELETE, PUT, PATCH",
		Headers: "*",
		MaxAge:  "64800",
	})

	// Auth API
	r.Post("/api/auth/sign-up", h.createUser)
	r.Post("/api/auth/login", h.login)
	r.Get("/api/auth/me", h.me)
	r.Get("/api/auth/check-auth", h.checkAuth)

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

func (h *Handler) handleAuth(r *http.Request) (string, error) {
	accessToken, err := extractToken(r)
	if err != nil {
		return "", errors.New("unauthorized error")
	}

	userId, err := h.service.Authorization.ParseAccessToken(accessToken)
	if err != nil {
		return "", errors.New("unauthorized error")
	}

	return userId, nil
}

// func (h *Handler) errorResponse(message string) map[string]string {
// 	res := make(map[string]string)
// 	res["status"] = "ERROR"
// 	res["message"] = message
// 	return res
// }

func (h *Handler) sendJSON(w http.ResponseWriter, payload interface{}, status int) {
	responseBody, err := json.Marshal(payload)

	if err != nil {
		logger.UseBasicLogger().Error("Internal server error", err, "func sendJSON()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(responseBody)
}

// func (h *Handler) parseJSONBody(w http.ResponseWriter, r *http.Request, pt interface{}) error {
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(pt)
// 	if err != nil {
// 		logger.UseBasicLogger().Error("Bad request", err, "func createGroup()")
// 		h.sendJSON(w, h.errorResponse("Cannot parse json"), http.StatusBadRequest)
// 		return err
// 	}

// 	return nil
// }
