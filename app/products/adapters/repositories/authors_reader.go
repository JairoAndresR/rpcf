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
	err := core.Injector.Provide(newAuthorsReader)
	core.CheckInjection(err, "newAuthorsReader")
}

type authorsReader struct {
	db *gorm.DB
}

func newAuthorsReader(conn sql.Connection) ports.AuthorsReader {
	return &authorsReader{db: conn.GetDatabase()}
}

func (r *authorsReader) GetAll() ([]*products.Author, error) {
	var authors []*entities.Author
	err := r.db.Find(&authors).Error
	return entities.MapAuthorsToDomain(authors), err
}

func (r *authorsReader) GetAllByGroup(groupCode string) ([]*products.Author, error) {
	var authors []*entities.Author
	err := r.db.
		Joins("join authors_gruplacs ag on ag.author_id = authors.id").
		Joins("join gruplacs g on g.id = ag.grup_lac_id and g.code =?", groupCode).
		Find(&authors).Error
	return entities.MapAuthorsToDomain(authors), err
}
