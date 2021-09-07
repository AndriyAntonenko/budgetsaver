package service

import (
	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(payload domain.User) (string, error) {
	hashedPassword, err := hashPassword(payload.Password)
	if err != nil {
		return "", err
	}

	return s.repo.CreateUser(domain.CreateUserRecord{
		Name:         payload.Name,
		Email:        payload.Email,
		PasswordHash: hashedPassword.passwordHash,
		Salt:         hashedPassword.salt,
	})
}
