package service

import (
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

type Service struct {
	Authorization
	FinanceGroup
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		FinanceGroup:  NewFinanceGroupService(repo.FinanceGroup),
	}
}
