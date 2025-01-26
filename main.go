package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	postgresRepo "github.com/hasbyadam/account-service/internal/repository/postgres"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/hasbyadam/account-service/account"
	"github.com/hasbyadam/account-service/internal/rest"
	"github.com/hasbyadam/account-service/internal/rest/middleware"

	"github.com/joho/godotenv"
)

const (
	defaultTimeout = 30
	defaultAddress = ":8081"
)

func init() {
	//log config
	log.SetFormatter(&log.JSONFormatter{})

	//prepare enviroment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Info("Success loading .env file")
}

func main() {
	//server config
	timeout := flag.Int("timeout", defaultTimeout, "timeout context (in seconds)")
	port := flag.String("port", defaultAddress, "server port")
	flag.Parse()
	log.Info("Starting server at port ", *port)

	//prepare database
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, "disable")
	dbConn, err := sql.Open(`postgres`, connection)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}
	log.Info("Success connect to database")

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("got error when closing the DB connection", err)
		}
	}()
	// prepare echo

	e := echo.New()
	e.Use(middleware.RequestLog)
	e.Use(middleware.CORS)

	timeoutContext := time.Duration(*timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// Prepare Repository
	accountRepo := postgresRepo.NewAccountRepository(dbConn)
	transactionRepo := postgresRepo.NewTransactionRepository(dbConn)

	// Build service Layer
	svc := account.NewService(transactionRepo, accountRepo)
	rest.NewAccountHandler(e, svc)

	// Start Server
	log.Fatal(e.Start(*port)) //nolint
}
