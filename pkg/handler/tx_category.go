package handler

import (
	"errors"
	"net/http"

	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	service "github.com/AndriyAntonenko/budgetSaver/pkg/services"
	"github.com/AndriyAntonenko/goRouter"
)

func (h *Handler) createTxCategory(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	logContext := "func createTxCategory()"
	userId, err := h.handleAuth(r)
	if err != nil {
		logger.UseBasicLogger().Error("Unauthorized error: ", err, logContext)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var txPayload dto.CreateTxCategoryDto
	err = h.parseJSONBody(r, &txPayload)
	if err != nil {
		logger.UseBasicLogger().Error("JSON parse error: ", err, logContext)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	txCategory, serviceErr := h.service.TxCategory.CreateTxCategory(userId, &txPayload)
	if serviceErr != nil {
		if serviceErr.Id == service.UnknownFinanceGroupMemberError {
			logger.UseBasicLogger().Error("Service error: ", errors.New(serviceErr.Error()), logContext)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		logger.UseBasicLogger().Error("Service error: ", errors.New(serviceErr.Error()), logContext)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, txCategory, http.StatusCreated)
}
