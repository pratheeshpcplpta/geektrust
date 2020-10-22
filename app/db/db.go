package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConn struct {
	DB       *gorm.DB
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	Domine   string
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func (db *DBConn) LoadConfig(configVal map[string]string) {
	db.Host = configVal["Host"]
	db.User = configVal["User"]
	db.Pass = configVal["Pass"]
	db.Port = configVal["Port"]
	db.Database = configVal["Database"]
	db.Domine = configVal["Domine"]
}

/**
 * -------------------------------------------------------------------------------
 * DB Connection
 * -------------------------------------------------------------------------------
 */
func (db *DBConn) DBConnection() {

	var connectiondb *gorm.DB
	var err error

	switch db.Domine {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", db.User, db.Pass, db.Host, db.Port, db.Database)
		connectiondb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			// Logger: logger.Default.LogMode(logger.Info),
		})
	case "sqlite":
		connectiondb, err = gorm.Open(sqlite.Open(db.Database+".db"), &gorm.Config{
			//Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		panic("Unable to find domine")
	}

	if err != nil {
		panic("failed to connect database")
	}
	db.DB = connectiondb
}

func (db *DBConn) Close() {
	db.Close()
}
