package ports

import "rpcf/accounts"

type TokenManager interface {
	Generate(acc *accounts.Account) (string, error)
	GetMetadata(tokenString string) (*accounts.Account, error)
	// It validates the to token
	Validate(token string) (bool, error)
}
