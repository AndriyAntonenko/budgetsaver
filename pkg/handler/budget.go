package handler

import (
	"net/http"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	"github.com/AndriyAntonenko/goRouter"
)

func (h *Handler) createBudget(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	userId, ok := h.getUserId(w, r)
	if !ok {
		return
	}

	var payload domain.CreateBudgetPayload
	err := h.parseJSONBody(w, r, &payload)
	if err != nil {
		return
	}

	budget, err := h.service.Budget.CreateBudget(userId, &payload)
	h.sendJSON(w, budget, http.StatusCreated)
}
