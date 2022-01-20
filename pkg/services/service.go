package service

import (
	"github.com/AndriyAntonenko/budgetSaver/pkg/config"
	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type Authorization interface {
	CreateUser(dto.UserSignUpPayload) (*Tokens, error)
	Login(dto.UserLoginPayload) (*Tokens, error)
	ParseRefreshToken(string) (string, error)
	ParseAccessToken(string) (string, error)
	GetProfile(string) (*dto.UserProfile, error)
}

type FinanceGroup interface {
	// pass user id and payload
	CreateFinanceGroup(string, dto.CreateFinanceGroupPayload) (*dto.FinanceGroup, error)
	// pass user id
	GetUsersFinanceGroups(string) ([]dto.FinanceGroup, error)
}

type Budget interface {
	// pass user id and payload
	CreateBudget(string, dto.CreateBudgetPayload) (*dto.Budget, *ServiceError)
	// user id, budget id
	FetchUserBudget(string, string) (*dto.Budget, *ServiceError)
}

type BudgetTx interface {
	// user id, budget id and payload
	CreateBudgetTx(string, string, dto.CreateBudgetTxDto) (*dto.BudgetTxDto, *ServiceError)
}

type TxCategory interface {
	// user id, payload
	CreateTxCategory(string, *dto.CreateTxCategoryDto) (*dto.TxCategoryDto, *ServiceError)
}

type CurracyExchange interface {
	GetSupportedSymbols() (map[string]string, error)
}

type Service struct {
	Authorization
	FinanceGroup
	Budget
	BudgetTx
	TxCategory
	CurracyExchange
}

func NewService(repo *repository.Repository) *Service {
	currencyExchangeService := NewCurracyExchangeService(config.UseAppConfig().ExchangeratesApi.ApiKey, config.UseAppConfig().ExchangeratesApi.Url)
	budgetService := NewBudgetService(repo.Budget, repo.FinanceGroup, currencyExchangeService)

	// TODO: Use pointers
	return &Service{
		Authorization:   NewAuthService(repo.Authorization),
		FinanceGroup:    NewFinanceGroupService(repo.FinanceGroup),
		Budget:          budgetService,
		BudgetTx:        NewTxService(&repo.BudgetTx, budgetService, &repo.TxCategory),
		TxCategory:      NewTxCategoryService(&repo.TxCategory, &repo.FinanceGroup),
		CurracyExchange: currencyExchangeService,
	}
}
