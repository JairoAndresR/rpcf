package sql

import "gorm.io/gorm"

// Connection allows to retrieve the gorm database connection.
type Connection interface {
	GetDatabase() *gorm.DB
}
