package main

import (
	"context"
	"fmt"
	db "github.com/avemoi/lacuba/db/sqlc"
	"github.com/gin-gonic/gin"
	"log"
)

func (app *Config) getLacubas(c *gin.Context) {
	res, err := app.Models.db.ListLacubas(context.Background())
	if err != nil {
		c.String(200, "We have an error ffs")
	}
	if res == nil {
		res = make([]db.Lacuba, 0)
	}
	c.JSON(200, res)
}

func (app *Config) addLacuba(c *gin.Context) {
	var newLacuba db.Lacuba
	if err := c.BindJSON(&newLacuba); err != nil {
		return
	}
	newlac, err := app.Models.db.CreateLacuba(context.Background(), db.CreateLacubaParams{
		Longtitude: newLacuba.Longtitude,
		Latitude:   newLacuba.Latitude,
	})
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println(newlac)
}
