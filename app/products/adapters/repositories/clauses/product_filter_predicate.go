package clauses

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"rpcf/app/products/adapters/repositories/entities"
	"sync"
)

// productFilterPredicate creates the SQL statement to be executed
type productFilterPredicate struct {
	db *gorm.DB
}

func NewProductFilterPredicate(db *gorm.DB) *productFilterPredicate {
	return &productFilterPredicate{db: db}
}

func (c *productFilterPredicate) Build(filters map[string]string) ([]string, string) {
	expressions := c.generateWhereExpression(filters)
	sql := c.buildFilter(expressions)
	values := c.generateValues(filters)
	return values, sql
}

func (c *productFilterPredicate) generateValues(filters map[string]string) []string {
	values := make([]string, 0)

	for _, v := range filters {
		values = append(values, v)
	}

	return values
}

func (c *productFilterPredicate) generateWhereExpression(filters map[string]string) []clause.Expression {
	expressions := make([]clause.Expression, 0)

	for column, v := range filters {
		value := fmt.Sprintf("%%%s%%", v)
		expr := clause.Like{
			Column: column,
			Value:  value,
		}
		expressions = append(expressions, expr)
	}
	return expressions
}

func (c *productFilterPredicate) buildFilter(expressions []clause.Expression) string {
	var buildNames []string
	var buildNamesMap = map[string]bool{}
	var clauses = c.generateClauses(expressions)

	stmt := c.generateStatement()
	for _, c := range clauses {
		if _, ok := buildNamesMap[c.Name()]; !ok {
			buildNames = append(buildNames, c.Name())
			buildNamesMap[c.Name()] = true
		}

		stmt.AddClause(c)
	}
	stmt.Build(buildNames...)
	return stmt.SQL.String()
}

func (c *productFilterPredicate) generateClauses(expressions []clause.Expression) []clause.Interface {

	if len(expressions) == 0 {
		return []clause.Interface{
			clause.Select{},
			clause.From{},
		}
	}

	return []clause.Interface{
		clause.Select{},
		clause.From{},
		clause.Where{Exprs: expressions},
	}
}

func (c *productFilterPredicate) generateStatement() gorm.Statement {
	productSchema, _ := schema.Parse(&entities.Product{}, &sync.Map{}, c.db.NamingStrategy)

	return gorm.Statement{
		DB:      c.db,
		Table:   productSchema.Table,
		Schema:  productSchema,
		Clauses: map[string]clause.Clause{},
	}
}
