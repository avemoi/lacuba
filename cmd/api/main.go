package main

import (
	"context"
	"database/sql"
	db "github.com/avemoi/lacuba/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

func openDB(cfg config) (*sql.DB, error) {
	conn, err := sql.Open("mysql", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = conn.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type repo struct {
	db *db.Queries
}

func NewRepo(db *db.Queries) *repo {
	return &repo{db: db}
}

func main() {
	var ginMode string
	if os.Getenv("release") == "true" {
		ginMode = gin.ReleaseMode
	} else {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)
	var cfg config
	cfg.env = "development"
	cfg.db.dsn = "root:mypassword@tcp(127.0.0.1:3307)/lacubadb?parseTime=true"

	conn, err := openDB(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	repo := NewRepo(db.New(conn))
	router := repo.GetRoutes()
	router.Run()

}
