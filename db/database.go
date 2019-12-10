package db

import (
	"github.com/jinzhu/gorm"
	//import Postgres Dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/samuelmjn/go-library/config"
)

// InitializeDatabase :nodoc:
func InitializeDatabase() (db *gorm.DB, err error) {
	dsn := config.DatabaseHost()
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic("cannot connect to database " + dsn)
	}
	return
}
