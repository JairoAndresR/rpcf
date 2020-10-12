package ports

import "rpcf/accounts"

type RegisterManager interface {
	Register(acc *accounts.Account) (*accounts.Account, string, error)
}
