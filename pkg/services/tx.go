package service

import (
	"database/sql"
	"time"

	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type TxService struct {
	txRepo         *repository.BudgetTx
	txCategoryRepo *repository.TxCategory
	budgetService  Budget
}

func NewTxService(tx *repository.BudgetTx, budget Budget, txCategoryRepo *repository.TxCategory) *TxService {
	return &TxService{
		txRepo:         tx,
		budgetService:  budget,
		txCategoryRepo: txCategoryRepo,
	}
}

func (s *TxService) CreateBudgetTx(userId string, budgetId string, payload dto.CreateBudgetTxDto) (*dto.BudgetTxDto, *ServiceError) {
	_, serviceErr := s.budgetService.FetchUserBudget(userId, budgetId)
	if serviceErr != nil {
		return nil, serviceErr
	}

	if payload.Category != nil {
		_, repoErr := (*s.txCategoryRepo).GetTxCategoryById(*payload.Category)
		if repoErr == sql.ErrNoRows {
			return nil, NewServiceError(EntityNotFound, "category not found")
		}
		if repoErr != nil {
			return nil, NewServiceError(UnexpectedError, repoErr.Error())
		}
	}

	newTx, repoErr := (*s.txRepo).CreateBudgetTx(repository.CreateBudgetTxRecord{
		BudgetId:    budgetId,
		Title:       payload.Title,
		Description: payload.Description,
		From:        payload.From,
		To:          payload.To,
		Amount:      payload.Amount,
		Author:      userId,
		Category:    payload.Category,
		TxTime:      time.Now(),
	})

	if repoErr != nil {
		return nil, NewServiceError(UnexpectedError, repoErr.Error())
	}

	res := dto.BudgetTxDto{
		Id:          newTx.Id,
		BudgetId:    newTx.BudgetId,
		Title:       newTx.Title,
		Description: newTx.Description,
		From:        newTx.From,
		To:          newTx.To,
		Amount:      newTx.Amount,
		Author:      newTx.Author,
		TxTime:      newTx.TxTime,
		Category:    &newTx.Category.String,
	}

	if newTx.Category.Valid {
		res.Category = &newTx.Category.String
	}

	return &res, nil
}
