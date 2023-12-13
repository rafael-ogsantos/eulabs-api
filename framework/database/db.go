package database

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rafael-ogsantos/eulabs-api/domain"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	dbInstance.Dsn = "root:root@tcp(db:3306)/eulabs?charset=utf8&parseTime=True&loc=Local"
	dbInstance.DsnTest = ":memory:"
	dbInstance.DbType = "mysql"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	db, err := dbInstance.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (d *Database) Connect() (*gorm.DB, error) {

	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       d.Dsn,
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	} else {
		d.Db, err = gorm.Open(sqlite.Open(d.DsnTest), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	}

	if err != nil {
		return nil, err
	}

	sqlDB, err := d.Db.DB()
	if err != nil {
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
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
