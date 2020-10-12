package ports

import "rpcf/accounts"

type LoginManager interface {
	Login(acc *accounts.Account) (*accounts.Account, string, error)
}
