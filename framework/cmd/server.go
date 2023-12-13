package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rafael-ogsantos/eulabs-api/application/repositories"
	"github.com/rafael-ogsantos/eulabs-api/application/services"
	"github.com/rafael-ogsantos/eulabs-api/framework/database"
	"github.com/rafael-ogsantos/eulabs-api/framework/router"
)

// Database instance
var db database.Database

// Load environment variables
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

// Main function
func main() {
	c, err := db.Connect()
	e := echo.New()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	productRepository := repositories.NewProductRepositoryDb(c)
	productService := services.NewProductService(productRepository)

	r := router.NewRouter(c, e, productService, productRepository)

	if err := r.Router().Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
