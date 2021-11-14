package handler

import (
	"net/http"

	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	"github.com/AndriyAntonenko/goRouter"
)

func (h *Handler) createFinanceGroup(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	userId, err := h.handleAuth(r)
	if err != nil {
		logger.UseBasicLogger().Error("Unauthorized error: ", err, "func createFinanceGroup()")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var groupPayload dto.CreateFinanceGroupPayload
	err = h.parseJSONBody(r, &groupPayload)
	if err != nil {
		logger.UseBasicLogger().Error("JSON parse error: ", err, "func createFinanceGroup()")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	financeGroup, err := h.service.FinanceGroup.CreateFinanceGroup(userId, groupPayload)
	if err != nil {
		logger.UseBasicLogger().Error("Service error: ", err, "func createFinanceGroup()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, financeGroup, http.StatusCreated)
}

func (h *Handler) fetchFinanceGroups(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	userId, err := h.handleAuth(r)
	if err != nil {
		logger.UseBasicLogger().Error("Unauthorized error: ", err, "func createFinanceGroup()")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	financeGroups, err := h.service.FinanceGroup.GetUsersFinanceGroups(userId)
	if err != nil {
		logger.UseBasicLogger().Error("Service error: ", err, "func createFinanceGroup()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, financeGroups, http.StatusOK)
}
