package initializers

import (
	"rpcf/app/dataproviders/sql"
	"rpcf/app/products/adapters/repositories/entities"
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

	db.AutoMigrate(entities.Articles{})
	db.AutoMigrate(entities.Books{})
	db.AutoMigrate(entities.ProductDefinition{})
	db.AutoMigrate(entities.Companies{})
	db.AutoMigrate(entities.DoctoralThesis{})
	db.AutoMigrate(entities.Product{})
}
