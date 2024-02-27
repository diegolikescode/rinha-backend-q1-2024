package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var Session *sql.DB

func SetupPostgres() {
    dbHost := "db"
    dbPort := "5432"
    dbName := "rinha"
    dbUser := "postgres"
    dbPassword := "postgres"

    connStr := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPassword)

    fmt.Println(connStr)
    fmt.Println()
    fmt.Println("CONN CONFIG:::", connStr)

    var err error
    Session, err = sql.Open("postgres", connStr)
    if err != nil {
	log.Fatal("ERROR CONNECTING TO POSTGRES", err)
    }

    Session.SetMaxOpenConns(10)
    Session.SetMaxIdleConns(2)
    Session.SetConnMaxLifetime(time.Minute * 5)

    log.Println("DATABASE CONFIGURED")
}

