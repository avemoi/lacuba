package main

import (
	"database/sql"
	db "github.com/avemoi/lacuba/db/sqlc"
	"log"
	"sync"
)

type models struct {
	db *db.Queries
}

func NewRepo(db *db.Queries) *models {
	return &models{db: db}
}

type Config struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	env      string
	Wait     *sync.WaitGroup
	Models   *models
}
