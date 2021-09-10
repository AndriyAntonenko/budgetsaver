package service

import (
	"errors"
	"time"

	"github.com/AndriyAntonenko/budgetSaver/pkg/config"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId string `json:"id"`
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

func (s *AuthService) ParseRefreshToken(refreshToken string) (string, error) {
	return parseToken(refreshToken, config.UseAppConfig().Jwt.RefreshTokenSecret)
}

func (s *AuthService) ParseAccessToken(accessToken string) (string, error) {
	return parseToken(accessToken, config.UseAppConfig().Jwt.AccessTokenSecret)
}

func parseToken(token string, secret string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok {
		return "", errors.New("token claims are not of type \"Claims\"")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return "", errors.New("token already expired")
	}

	return claims.UserId, nil
}

func generateTokens(id string) (*Tokens, error) {
	accessExpirationDate := time.Now().Add(accessExpiresIn).Unix()
	refreshExpirationDate := time.Now().Add(refreshExpiresIn).Unix()

	accessClaims := &Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationDate,
		},
	}

	refreshClaims := &Claims{
		UserId: id,
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
