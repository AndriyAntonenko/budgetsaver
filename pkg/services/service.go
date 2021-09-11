package service

import (
	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type Authorization interface {
	CreateUser(domain.UserSignUpPayload) (string, error)
	Login(domain.UserLoginPayload) (*Tokens, error)
	ParseRefreshToken(string) (string, error)
	ParseAccessToken(string) (string, error)
	GetProfile(string) (*domain.UserProfile, error)
}

type Group interface {
	CreateGroup(string, *domain.CreateGroupPayload) (*domain.Group, error)
}

type Service struct {
	Authorization
	Group
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Group:         NewGroupService(repo.Group),
	}
}
