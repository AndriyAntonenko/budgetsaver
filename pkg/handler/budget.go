package handler

import (
	"errors"
	"net/http"

	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	service "github.com/AndriyAntonenko/budgetSaver/pkg/services"
	"github.com/AndriyAntonenko/goRouter"
)

func (h *Handler) createBudget(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	userId, err := h.handleAuth(r)
	if err != nil {
		logger.UseBasicLogger().Error("Unauthorized error: ", err, "func createBudget()")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var budgetPayload dto.CreateBudgetPayload
	err = h.parseJSONBody(r, &budgetPayload)
	if err != nil {
		logger.UseBasicLogger().Error("JSON parse error: ", err, "func createBudget()")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	budget, serviceErr := h.service.Budget.CreateBudget(userId, budgetPayload)
	if serviceErr != nil {
		if serviceErr.Id == service.ActionForbiddenError || serviceErr.Id == service.UnknownFinanceGroupMemberError {
			logger.UseBasicLogger().Error("Service error: ", errors.New(serviceErr.Error()), "func createBudget()")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		logger.UseBasicLogger().Error("Service error: ", errors.New(serviceErr.Error()), "func createBudget()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, budget, http.StatusCreated)
}
