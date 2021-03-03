package repositories

import (
	"fmt"
	"rpcf/app/dataproviders/sql"
	"rpcf/app/products/adapters/repositories/entities"
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"

	"gorm.io/gorm"
)

func init() {
	err := core.Injector.Provide(newReportsReader)
	core.CheckInjection(err, "newReportsReader")
}

type CountResult struct {
	Value string
	Count int64
}

type ReportsReader struct {
	db *gorm.DB
}

func newReportsReader(conn sql.Connection) ports.ReportsReader {
	return &ReportsReader{db: conn.GetDatabase()}
}

func (r *ReportsReader) CountAllBy(filters map[string]string, groupType string) ([]*products.Report, error) {
	results := make([]*products.Report, 0)

	selectQuery := fmt.Sprintf("%s as value, COUNT(%s) as count", groupType, groupType)
	tx := r.db.Model(entities.Product{}).Select(selectQuery)
	for column, value := range filters {
		q := fmt.Sprintf(`%s LIKE ?`, column)
		tx.Where(q, fmt.Sprintf("%%%s%%", value))
	}

	err := tx.Group(groupType).Find(&results).Error
	return results, err
}

func (r *ReportsReader) CountProductsByCategory() []*products.Report {

	categories := map[string][]string{
		"applied_investigation": {"master_theses", "tech_products", "software", "articles", "inv_books"},
		"pure_investigation":    {"doctoral_theses"},
		"formation":             {"text_books", "books_traduction", "theses", "presentations", "manuals"},
		"consulting":            {"startup", "spinoff", "consultancies"},
	}

	results := make([]*products.Report, 0)

	for category, ps := range categories {
		var count int64

		err := r.db.Model(entities.Product{}).Where("type_name IN ?", ps).Count(&count).Error

		if err != nil {
			return nil
		}

		r := products.Report{
			Value: category,
			Count: count,
		}

		results = append(results, &r)
	}

	return results
}
