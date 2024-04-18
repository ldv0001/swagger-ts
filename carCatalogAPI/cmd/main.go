package main

import (
	"catalog/handler"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "catalog/docs"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Car Catalog API
// @version 1.0
// @description This is a sample server Car Catalog server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	sourceURL := os.Getenv("sourceURL")
	databaseURL := os.Getenv("databaseURL")

	m, err := migrate.New(
		sourceURL,
		databaseURL)

	if err != nil {
		log.Println(err)
	}
	if err := m.Up(); err != nil {
		log.Println("migration log: ", err)
	}
	m.Close()

}
func main() {
	log.Println("server started")
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	port := os.Getenv("localport")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.GetCars)
	mux.HandleFunc("/add", handler.AddCars)
	mux.HandleFunc("/delete", handler.Delete)
	mux.HandleFunc("/update", handler.Update)
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.ListenAndServe(port, mux)

}
