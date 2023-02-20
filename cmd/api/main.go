package main

import (
	"database/sql"
	"fmt"
	db "github.com/avemoi/lacuba/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	connDB := initDB()
	defer connDB.Close()

	// create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	wg := sync.WaitGroup{}

	// set up the application config
	app := Config{
		DB:       connDB,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Wait:     &wg,
	}

	if os.Getenv("release") == "true" {
		app.env = gin.ReleaseMode
	} else {
		app.env = gin.DebugMode
	}
	gin.SetMode(app.env)

	app.Models = NewRepo(db.New(connDB))

	router := app.GetRoutes()
	router.Run(fmt.Sprintf(":%s", os.Getenv("GINPORT")))

}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can not connect to database")
	}
	return conn
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("mysql not yet ready")
		} else {
			log.Println("Connected to datbaase")
			return connection
		}

		if counts > 10 {
			return nil
		}
		log.Println("backing off for 1 seconds")
		time.Sleep(1 * time.Second)
		counts++
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	fmt.Println("trying wiht", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil

}
