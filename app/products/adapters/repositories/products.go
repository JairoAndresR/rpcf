package repositories

import "rpcf/app/products/adapters/repositories/entities"

var products = map[string]interface{}{
	"articles": &entities.Articles{},
	"books":    &entities.Books{},
}
