package initializer

import (
	"rpcf/app/accounts/adapters/repositories/entities"
	"rpcf/app/dataproviders/sql"
	"rpcf/core"
)

func Migrate() {
	var conn sql.Connection
	invokeFunc := func(connection sql.Connection) {
		conn = connection
	}

	err := core.Injector.Invoke(invokeFunc)
	if err != nil {
		panic(err)
	}

	db := conn.GetDatabase()

	db.AutoMigrate(entities.Account{})
}
