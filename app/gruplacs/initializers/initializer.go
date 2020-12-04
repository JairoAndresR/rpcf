package initializers

import (
	"rpcf/app/dataproviders/sql"
	"rpcf/app/gruplacs/adapters/repositories/entities"
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

	db.AutoMigrate(&entities.GruplacDefinition{})
}
