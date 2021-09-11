package handler

import (
	"net/http"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	"github.com/AndriyAntonenko/goRouter"
)

func (h *Handler) createGroup(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	userId, ok := h.getUserId(w, r)
	if !ok {
		return
	}

	var payload domain.CreateGroupPayload
	h.parseJSONBody(w, r, &payload)

	group, err := h.service.Group.CreateGroup(userId, &payload)
	if err != nil {
		logger.UseBasicLogger().Error("Internal server error", err, "func createGroup()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, group, http.StatusCreated)
}
