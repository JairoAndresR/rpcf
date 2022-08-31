package ports

import (
	"encoding/json"
	"fmt"
	"rpcf/products"
	"strings"
)

// Represents the generic information of any kind of product in the system
type GenericProduct interface {
	GetID() string
}

func NewGenericProduct(generic interface{}, p *products.ProductResult) *products.Product {
	for key, value := range p.Fields {
		fmt.Println("KEY:", key, value)
	}
	b, err := json.Marshal(generic)

	if err != nil {
		return nil
	}

	var product products.Product
	err = json.Unmarshal(b, &product)

	if err != nil {
		return nil
	}
	product.TypeId = product.ID
	product.TypeName = p.Name
	product.GrouplacCode = p.GrupLACCode
	product.GroupName = p.GrupLACName
	product.Authors = p.Authors
	fmt.Println("TYPEEE")
	fmt.Println(p.Fields["type"])
	if p.Fields["type"] == "Publicado en revista especializada" {
		product.StartYear = p.Fields["published_year"]
	}
	if strings.Contains(p.Fields["type"], "Libro resultado de") {
		product.StartYear = p.Fields["published_year"]
	}
	return &product
}
