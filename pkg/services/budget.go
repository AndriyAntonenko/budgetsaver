package service

import (
	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type BudgetService struct {
	budgetRepo repository.Budget
}

func NewBudgetService(budgetRepo repository.Budget) *BudgetService {
	return &BudgetService{budgetRepo}
}

func (bs *BudgetService) CreateBudget(userId string, payload *domain.CreateBudgetPayload) (*domain.Budget, error) {
	return bs.budgetRepo.CreateBudget(userId, payload)
}
