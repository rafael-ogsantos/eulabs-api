package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rafael-ogsantos/eulabs-api/framework/database"
	"github.com/rafael-ogsantos/eulabs-api/framework/router"
)

var db database.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	autoMigrateDb, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE_DB"))
	if err != nil {
		log.Fatal("Error parsing AUTO_MIGRATE_DB")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatal("Error parsing DEBUG")
	}

	db.AutoMigrateDb = autoMigrateDb
	db.Debug = debug
	db.DsnTest = os.Getenv("DSN_TEST")
	db.Dsn = os.Getenv("DSN")
	db.DbType = os.Getenv("DB_TYPE")
	db.DbTypeTest = os.Getenv("DB_TYPE_TEST")
	db.Env = os.Getenv("ENV")
}

func main() {
	c, err := db.Connect()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	e := router.New(c)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
