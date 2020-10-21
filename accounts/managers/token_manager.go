package managers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"rpcf/accounts"
	"rpcf/accounts/ports"
	"rpcf/core"
	"time"
)

const (
	expirationTime    = time.Minute * 90
	tokenAccessSecret = "TOKEN_ACCESS_SECRET"
)

func init() {
	err := core.Injector.Provide(newTokenManager)
	core.CheckInjection(err, "TokenManager")
}

type tokenManager struct {
}

func newTokenManager() ports.TokenManager {
	return &tokenManager{}
}

func (m *tokenManager) Generate(acc *accounts.Account) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true

	claims["account_id"] = acc.ID
	claims["names"] = acc.Names
	claims["email"] = acc.Email
	claims["exp"] = time.Now().Add(expirationTime).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenAccessSecret := os.Getenv(tokenAccessSecret)
	sign := []byte(tokenAccessSecret)

	token, err := at.SignedString(sign)
	if err != nil {
		return "", errors.New(TokenError)
	}
	return token, nil
}

func (m *tokenManager) getJWTToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header)
		}
		return []byte(os.Getenv("TOKEN_ACCESS_SECRET")), nil
	})

	return token, err
}

func (m *tokenManager) GetMetadata(tokenString string) (*accounts.Account, error) {
	token, err := m.getJWTToken(tokenString)

	if err != nil {
		return nil, errors.New("invalid token error")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	accountID := claims["account_id"]
	names := claims["names"]
	email := claims["email"]
	claims["exp"] = time.Now().Add(expirationTime).Unix()

	user := &accounts.Account{
		ID:    (accountID).(string),
		Names: (names).(string),
		Email: (email).(string),
	}

	return user, nil
}

func (m *tokenManager) Validate(tokenString string) (bool, error) {
	token, err := m.getJWTToken(tokenString)

	if err != nil {
		return false, errors.New(InvalidTokenError)
	}
	return token.Valid, nil
}
