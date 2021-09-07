package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	"github.com/AndriyAntonenko/goRouter"
)

// TODO: Finish this method
func (h *Handler) createUser(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	var payload domain.User
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
