package service

import (
	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type BudgetService struct {
	budgetRepo              repository.Budget
	financeGroupRepo        repository.FinanceGroup
	currencyExchangeService CurracyExchange
}

func NewBudgetService(budgetRepo repository.Budget, financeGroupRepo repository.FinanceGroup, currencyExchangeService CurracyExchange) *BudgetService {
	return &BudgetService{
		budgetRepo:              budgetRepo,
		financeGroupRepo:        financeGroupRepo,
		currencyExchangeService: currencyExchangeService,
	}
}

func (s *BudgetService) CreateBudget(userId string, payload dto.CreateBudgetPayload) (*dto.Budget, *ServiceError) {
	role, err := s.financeGroupRepo.GetUserRoleInFinanceGroup(payload.GroupId, userId)

	if err != nil {
		return nil, NewServiceError(UnexpectedError, err.Error())
	}

	if role == "" {
		return nil, NewServiceError(UnknownFinanceGroupMemberError, "user not in this group")
	}

	if role == domain.Member {
		return nil, NewServiceError(ActionForbiddenError, "user should be administrator or owner of this group to create a budget")
	}

	availableSymbols, err := s.currencyExchangeService.GetSupportedSymbols()

	if err != nil {
		return nil, NewServiceError(UnexpectedError, "cannot verify currency")
	}

	if _, currencyExists := availableSymbols[payload.Currency]; !currencyExists {
		return nil, NewServiceError(WrongPropertyValues, "such currency not exists")
	}

	newBudget, err := s.budgetRepo.CreateBudget(repository.CreateBudgetRecord{
		Creator:        userId,
		FinanceGroupId: payload.GroupId,
		Name:           payload.Name,
		Description:    payload.Description,
		Currency:       payload.Currency,
	})

	if err != nil {
		return nil, NewServiceError(UnexpectedError, err.Error())
	}

	return &dto.Budget{
		Id:          newBudget.Id,
		Name:        newBudget.Name,
		Description: newBudget.Description,
		Creator:     newBudget.Creator,
		GroupId:     newBudget.FinanceGroupId,
		Currency:    newBudget.Currency,
	}, nil
}

func (s *BudgetService) FetchUserBudget(userId string, budgetId string) (*dto.Budget, *ServiceError) {
	budget, err := s.budgetRepo.GetUserBudget(budgetId, userId)
	if err != nil {
		return nil, NewServiceError(EntityNotFound, err.Error())
	}

	return &dto.Budget{
		Id:          budget.Id,
		Name:        budget.Name,
		Description: budget.Description,
		Creator:     budget.Creator,
		GroupId:     budget.FinanceGroupId,
		Currency:    budget.Currency,
	}, nil
}
