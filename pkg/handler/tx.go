package handler

import (
	"errors"
	"net/http"

	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	service "github.com/AndriyAntonenko/budgetSaver/pkg/services"
	"github.com/AndriyAntonenko/goRouter"
)

func (h *Handler) createTx(w http.ResponseWriter, r *http.Request, rp *goRouter.RouterParams) {
	logContext := "func createTx()"
	userId, err := h.handleAuth(r)
	if err != nil {
		logger.UseBasicLogger().Error("Unauthorized error: ", err, logContext)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var txPayload dto.CreateBudgetTxDto
	err = h.parseJSONBody(r, &txPayload)
	if err != nil {
		logger.UseBasicLogger().Error("JSON parse error: ", err, logContext)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	budgetId := rp.GetString("budgetId")
	newTx, serviceErr := h.service.BudgetTx.CreateBudgetTx(userId, budgetId, txPayload)
	if serviceErr != nil {
		if serviceErr.Id == service.EntityNotFound {
			logger.UseBasicLogger().Error("Service error: ", errors.New(serviceErr.Error()), logContext)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger.UseBasicLogger().Error("Service error: ", errors.New(serviceErr.Error()), logContext)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, newTx, http.StatusCreated)
}
