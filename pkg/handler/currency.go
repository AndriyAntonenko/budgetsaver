package handler

import (
	"net/http"

	"github.com/AndriyAntonenko/budgetSaver/pkg/logger"
	"github.com/AndriyAntonenko/goRouter"
)

func (h *Handler) getAvailableCurrencies(w http.ResponseWriter, r *http.Request, _ *goRouter.RouterParams) {
	logContext := "func getAvailableCurrencies()"
	symbols, err := h.service.CurracyExchange.GetSupportedSymbols()

	if err != nil {
		logger.UseBasicLogger().Error("Service error: ", err, logContext)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, symbols, http.StatusCreated)
}
