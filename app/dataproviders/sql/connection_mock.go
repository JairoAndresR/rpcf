package sql

import (
	"database/sql"
	mocket "github.com/selvatico/go-mocket"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MockConnection is an usefull connection for unit testing
type MockConnection struct {
	db *gorm.DB
}

// NewMockConnection setup a mock connection.
func NewMockConnection() Connection {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	sqlDB, _ := sql.Open(mocket.DriverName, "connection_mock")
	db, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	return &MockConnection{db: db}
}

// GetDatabase retrieves a mock connection.
func (c *MockConnection) GetDatabase() *gorm.DB {
	return c.db
}
