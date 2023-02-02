package main

import (
	"context"
	"fmt"
	db "github.com/avemoi/lacuba/db/sqlc"
	"github.com/gin-gonic/gin"
	"log"
)

func (r *repo) getLacubas(c *gin.Context) {

	res, err := r.db.ListLacubas(context.Background())
	if err != nil {
		c.String(200, "We have an error ffs")
	}
	if res == nil {
		res = make([]db.Lacuba, 0)
	}
	c.JSON(200, res)
}

func (r *repo) addLacuba(c *gin.Context) {
	var newLacuba db.Lacuba
	if err := c.BindJSON(&newLacuba); err != nil {
		return
	}
	newlac, err := r.db.CreateLacuba(context.Background(), db.CreateLacubaParams{
		Longtitude: newLacuba.Longtitude,
		Latitude:   newLacuba.Latitude,
	})
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println(newlac)
}
