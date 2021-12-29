package service

import (
	"time"

	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type TxService struct {
	txRepo        repository.BudgetTx
	budgetService Budget
}

func NewTxService(tx repository.BudgetTx, budget Budget) *TxService {
	return &TxService{
		txRepo:        tx,
		budgetService: budget,
	}
}

func (s *TxService) CreateBudgetTx(userId string, budgetId string, payload dto.CreateBudgetTxDto) (*dto.BudgetTxDto, *ServiceError) {
	_, err := s.budgetService.FetchUserBudget(userId, budgetId)
	if err != nil {
		return nil, err
	}

	newTx, repoErr := s.txRepo.CreateBudgetTx(repository.CreateBudgetTxRecord{
		BudgetId:    budgetId,
		Title:       payload.Title,
		Description: payload.Description,
		From:        payload.From,
		To:          payload.To,
		Amount:      payload.Amount,
		Author:      userId,
		TxTime:      time.Now(),
	})

	if repoErr != nil {
		return nil, NewServiceError(UnexpectedError, repoErr.Error())
	}

	return &dto.BudgetTxDto{
		Id:          newTx.Id,
		BudgetId:    newTx.BudgetId,
		Title:       newTx.Title,
		Description: newTx.Description,
		From:        newTx.From,
		To:          newTx.To,
		Amount:      newTx.Amount,
		Author:      newTx.Author,
		TxTime:      newTx.TxTime,
	}, nil
}
