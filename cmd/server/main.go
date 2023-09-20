package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/marinazv/FinalGo/cmd/server/routes"
	_ "github.com/marinazv/FinalGo/docs"
	"github.com/marinazv/FinalGo/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Examen Final Backend III
// @version         1.0
// @description     Desafío Final: Sistema de reserva de turnos
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9090
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	// Recover from panic.
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Load the environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database.
	db := connectDB()

	// Create a new Gin engine.
	engine := gin.New()
	engine.Use(gin.Recovery())
	// Add the logger middleware.
	engine.Use(middleware.Logger())

	// Add the swagger handler.
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the application.
	runApp(db, engine)

	// Close the connection.
	defer db.Close()

}

func runApp(db *sql.DB, engine *gin.Engine) {
	// Run the application.
	router := routes.NewRouter(engine, db)
	// Map all routes.
	router.MapRoutes()
	if err := engine.Run(os.Getenv("LOCAL_PORT")); err != nil {
		panic(err)
	}

}

// connectDB connects to the database.
func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

	// Create the data source.
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	fmt.Print(dataSource)

	// Open the connection.
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	// Check the connection.
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	// Leer el contenido del archivo
	sqlBytes, err := os.ReadFile("schema.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convertir el contenido a una cadena
	sqlString := string(sqlBytes)

	// Divide el contenido en sentencias SQL individuales
	sqlStatements := strings.Split(sqlString, ";")

	// Ejecutar cada sentencia SQL en la base de datos
	for _, statement := range sqlStatements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}

		if _, err := db.Exec(statement); err != nil {
			log.Printf("Error al ejecutar sentencia SQL: %v", err)
		}
	}
	return db
}
