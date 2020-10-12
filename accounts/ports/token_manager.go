package ports

import "rpcf/accounts"

type TokenManager interface {
	Generate(acc *accounts.Account) (string, error)
}
