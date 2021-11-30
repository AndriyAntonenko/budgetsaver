package service

import (
	"errors"
	"time"

	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(payload dto.UserSignUpPayload) (*Tokens, error) {
	hashedPassword, err := hashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	iat := time.Now()
	userId, err := s.repo.CreateUser(repository.CreateUserRecord{
		Name:         payload.Name,
		Email:        payload.Email,
		PasswordHash: hashedPassword.passwordHash,
		Salt:         hashedPassword.salt,
		LastLoginAt:  iat,
	})

	if err != nil {
		return nil, err
	}

	return generateTokens(userId, iat)
}

func (s *AuthService) Login(payload dto.UserLoginPayload) (*Tokens, error) {
	user, err := s.repo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	hashString, err := hashPasswordWithSalt(payload.Password, user.Salt)
	if err != nil {
		return nil, err
	}

	if hashString != user.PasswordHash {
		return nil, errors.New("password is not match")
	}

	iat := time.Now()
	user, err = s.repo.UpdateLastLogin(user.Id, iat)
	if err != nil {
		return nil, err
	}

	return generateTokens(user.Id, iat)
}

func (s *AuthService) GetProfile(id string) (*dto.UserProfile, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserProfile{
		UserId: user.Id,
		Email:  user.Email,
		Name:   user.Name,
	}, nil
}
