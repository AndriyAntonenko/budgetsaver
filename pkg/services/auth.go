package service

import (
	"errors"

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

	userId, err := s.repo.CreateUser(repository.CreateUserRecord{
		Name:         payload.Name,
		Email:        payload.Email,
		PasswordHash: hashedPassword.passwordHash,
		Salt:         hashedPassword.salt,
	})

	if err != nil {
		return nil, err
	}

	return generateTokens(userId)
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

	return generateTokens(user.Id)
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
