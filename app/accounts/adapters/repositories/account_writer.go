package repositories

import (
	"github.com/jinzhu/gorm"
	"rpcf/accounts"
	"rpcf/accounts/ports"
	"rpcf/app/accounts/adapters/repositories/entities"
	"rpcf/app/dataproviders/sql"
	"rpcf/core"
)

func init() {
	err := core.Injector.Provide(newAccountWriter)
	core.CheckInjection(err, "AccountReader")
}

type accountWriter struct {
	db *gorm.DB
}

func newAccountWriter(conn sql.Connection) ports.AccountWriter {
	db := conn.GetDatabase()
	return &accountWriter{db}
}

func (u *accountWriter) Create(account *accounts.Account) (*accounts.Account, error) {
	entity := entities.NewFromDomain(account)
	err := u.db.Model(entities.Account{}).Save(&entity).Error

	if err != nil {
		return nil, err
	}
	return entity.ToDomain(), nil
}
