package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	"github.com/AndriyAntonenko/goRouter"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	var payload domain.UserSignUp
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&payload)
	if err != nil {
		logger.UseBasicLogger().Error("Bad request", err, "func createUser()")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateUser(payload)
	if err != nil {
		logger.UseBasicLogger().Error("Internal server error", err, "func createUser()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(map[string]interface{}{
		"id": id,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	var payload domain.UserLoginPayload
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

	responseBody, err := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}
