package service

import (
	"errors"

	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(payload domain.UserSignUpPayload) (string, error) {
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

func (s *AuthService) Login(payload domain.UserLoginPayload) (*Tokens, error) {
	user, err := s.repo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	hashString, err := hashPasswordWithSalt(payload.Password, user.Salt)
	if err != nil {
		return nil, err
	}

	if hashString != user.PasswordHash {
		return nil, errors.New("Password is not match")
	}

	return generateTokens(user.Id)
}
