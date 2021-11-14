package handler

import (
	"encoding/json"
	"net/http"

	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	"github.com/AndriyAntonenko/goRouter"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	var payload dto.UserSignUpPayload
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&payload)
	if err != nil {
		logger.UseBasicLogger().Error("Bad request", err, "func createUser()")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokens, err := h.service.CreateUser(payload)
	if err != nil {
		logger.UseBasicLogger().Error("Internal server error", err, "func createUser()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, tokens, http.StatusCreated)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	var payload dto.UserLoginPayload
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&payload)
	if err != nil {
		logger.UseBasicLogger().Error("Bad request", err, "func login()")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := h.service.Login(payload)
	if err != nil {
		logger.UseBasicLogger().Error("Internal server error", err, "func createUser()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, data, http.StatusOK)
}

func (h *Handler) me(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	userId, err := h.handleAuth(r)
	if err != nil {
		logger.UseBasicLogger().Error("Unauthorized error", err, "func me()")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	profile, err := h.service.Authorization.GetProfile(userId)
	if err != nil {
		logger.UseBasicLogger().Error("Internal server error", err, "func createUser()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, profile, http.StatusOK)
}

func (h *Handler) checkAuth(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	_, err := h.handleAuth(r)
	if err != nil {
		logger.UseBasicLogger().Error("Unauthorized error", err, "func checkAuth()")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}
