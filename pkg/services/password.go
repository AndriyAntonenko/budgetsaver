package service

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type HashedPassword struct {
	passwordHash string
	salt         string
}

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateSalt() string {
	b := make([]rune, PW_SALT_BYTES)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func hashPassword(password string) (*HashedPassword, error) {
	salt := generateSalt()
	hash := sha1.New()

	_, err := hash.Write([]byte(password))
	if err != nil {
		return nil, err
	}

	return &HashedPassword{
		passwordHash: fmt.Sprintf("%x", hash.Sum([]byte(salt))),
		salt:         salt,
	}, nil
}
