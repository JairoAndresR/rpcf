package repositories

import (
	"gorm.io/gorm"
	"rpcf/app/dataproviders/sql"
	"rpcf/app/products/adapters/repositories/entities"
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newAuthorsWriter)
	core.CheckInjection(err, "newAuthorsWriter")
}

type authorsWriter struct {
	db *gorm.DB
}

func newAuthorsWriter(conn sql.Connection) ports.AuthorsWriter {
	return &authorsWriter{db: conn.GetDatabase()}
}

func (w *authorsWriter) Create(a *products.Author) (*products.Author, error) {
	author := entities.NewAuthor(a)
	err := w.db.Create(&author).Error
	if err != nil {
		return nil, err
	}
	return author.ToDomain(), nil
}
