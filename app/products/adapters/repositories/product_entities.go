package repositories

import "rpcf/app/products/adapters/repositories/entities"

var productEntities = map[string]interface{}{
	"articles": &entities.Articles{},
	"books":    &entities.Books{},
}
