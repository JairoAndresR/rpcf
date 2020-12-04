package repositories

import (
	"gorm.io/gorm"
	"rpcf/accounts"
	"rpcf/accounts/ports"
	"rpcf/app/accounts/adapters/repositories/entities"
	"rpcf/app/dataproviders/sql"
	"rpcf/core"
)

func init() {
	err := core.Injector.Provide(newAccountReader)
	core.CheckInjection(err, "AccountReader")
}

type accountReader struct {
	db *gorm.DB
}

func newAccountReader(conn sql.Connection) ports.AccountReader {
	db := conn.GetDatabase()
	return &accountReader{db}
}

func (u *accountReader) FindByID(id string) *accounts.Account {
	var account entities.Account
	u.db.Where("id = ?", id).First(&account)
	return account.ToDomain()
}

func (u *accountReader) FindByEmail(email string) *accounts.Account {
	var account entities.Account
	u.db.Where("email = ?", email).First(&account)
	return account.ToDomain()
}
