package repositories

import "rpcf/app/products/adapters/repositories/entities"

var productEntities = map[string]interface{}{
	entities.ArticlesTableName: entities.Articles{},
	"books":                    entities.Books{},
}
