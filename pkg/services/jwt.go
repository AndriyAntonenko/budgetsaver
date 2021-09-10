package service

import (
	"time"

	"github.com/AndriyAntonenko/budgetSaver/pkg/config"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken           string `json:"accessToken"`
	RefreshToken          string `json:"refreshToken"`
	AccessTokenExpiresAt  int64  `json:"accessTokenExpiresAt"`  // unix time
	RefreshTokenExpiresAt int64  `json:"refreshTokenExpiresAt"` // unix time
}

const (
	accessExpiresIn  = time.Hour * 12
	refreshExpiresIn = time.Hour * 24 * 3
)

func generateTokens(id string) (*Tokens, error) {
	accessExpirationDate := time.Now().Add(accessExpiresIn).Unix()
	refreshExpirationDate := time.Now().Add(refreshExpiresIn).Unix()

	accessClaims := &Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationDate,
		},
	}

	refreshClaims := &Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationDate,
		},
	}

	accessSecret := []byte(config.UseAppConfig().Jwt.AccessTokenSecret)
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(accessSecret)
	if err != nil {
		return nil, err
	}

	refreshSecret := []byte(config.UseAppConfig().Jwt.RefreshTokenSecret)
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(refreshSecret)
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessExpirationDate,
		RefreshTokenExpiresAt: refreshExpirationDate,
	}, nil
}
