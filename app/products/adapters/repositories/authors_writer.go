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
	//here we are going to reset the database
	author := entities.NewAuthor(a)
	err := w.db.Create(&author).Error
	if err != nil {
		return nil, err
	}
	return author.ToDomain(), nil
}

func (w *authorsWriter) ClearDB() {
	w.db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	w.db.Exec("TRUNCATE TABLE authors_products;")
	w.db.Exec("TRUNCATE TABLE authors_gruplacs;")
	w.db.Exec("TRUNCATE TABLE products;")
	w.db.Exec("TRUNCATE TABLE authors;")
	w.db.Exec("TRUNCATE TABLE articles;")
	w.db.Exec("TRUNCATE TABLE books;")
	w.db.Exec("TRUNCATE TABLE companies;")
	w.db.Exec("TRUNCATE TABLE doctoral_theses;")
	w.db.Exec("SET FOREIGN_KEY_CHECKS = 1;")	
}
