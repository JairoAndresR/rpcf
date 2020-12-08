package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockConnection(t *testing.T) {
	conn := NewMockConnection()
	db := conn.GetDatabase()

	sqlDB, err := db.DB()
	assert.Nil(t, err)
	isAlive := sqlDB.Ping()
	assert.Nil(t, isAlive)
}
