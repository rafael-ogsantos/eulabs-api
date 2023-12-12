package database

import (
	"log"

	"github.com/rafael-ogsantos/eulabs-api/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDatabase() *Database {
	return &Database{}
}

func NewDatabaseTest() *gorm.DB {
	dbInstance := NewDatabase()
	dbInstance.Env = "test"
	dbInstance.DbType = "sqlite3"
	dbInstance.DsnTest = ":memory"
	dbInstance.AutoMigrateDb = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {

	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(mysql.Open(d.Dsn))
	} else {
		d.Db, err = gorm.Open(mysql.Open(d.DsnTest))
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db = d.Db.Debug()
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Product{})
	}

	return d.Db, nil
}
