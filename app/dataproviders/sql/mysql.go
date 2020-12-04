package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"rpcf/core"
	"time"

	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := core.Injector.Provide(NewMySQLConnection)
	core.CheckInjection(err, "MySQLConnection")
}

// getURL retrieves the URL to connection to SQL database.
func getURL(params ...string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", params[0], params[1], params[2], params[3], params[4])
}

// MySQLConnection represents a connection for MySQL database
type MySQLConnection struct {
	db *gorm.DB
}

// NewMySQLConnection retrieves a sql connection to MySQL server
func NewMySQLConnection() (Connection, error) {
	dbUsername := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_USER_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := getURL(dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	log.Println(url)

	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		return nil, err
	}

	return &MySQLConnection{db: db}, nil
}

// GetDatabase retrieve a MySQL connection
func (c *MySQLConnection) GetDatabase() *gorm.DB {
	return c.db
}
