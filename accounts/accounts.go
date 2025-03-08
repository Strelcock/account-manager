package accounts

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID          uuid.UUID
	Name        string
	Email       string
	Password    string
	TimeCreated time.Time
	TimeUpdated time.Time
}

const defaultPasswordLength = 12

var symbols = []string{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*?,.[]{}"}

func generatePassword(length int) string {
	password := ""
	for range length {
		max := big.NewInt(int64(len(symbols)))
		random, _ := rand.Int(rand.Reader, max)
		password += symbols[int(random.Int64())]
	}
	return password

}

func NewAccount(name, email, password string) (*Account, error) {
	if strings.TrimSpace(name) == "" {
		return nil, NewAccountError("empty name", emptyName)
	}
	if strings.TrimSpace(email) == "" {
		return nil, NewAccountError("empty email", emptyEmail)
	}

	if strings.TrimSpace(password) == "" {
		password = generatePassword(defaultPasswordLength)
	}

	return &Account{
		ID:          uuid.New(),
		Name:        name,
		Email:       email,
		Password:    password,
		TimeCreated: time.Now(),
		TimeUpdated: time.Now(),
	}, nil
}
