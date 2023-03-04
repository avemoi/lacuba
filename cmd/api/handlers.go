package main

import (
	"context"
	"fmt"
	db "github.com/avemoi/lacuba/db/sqlc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	lacubaResult, err := app.Models.db.CreateLacuba(context.Background(), db.CreateLacubaParams{
		Longtitude: newLacuba.Longtitude,
		Latitude:   newLacuba.Latitude,
	})
	if err != nil {
		log.Fatal("error", err)
	}

	newLacubaId, err := lacubaResult.LastInsertId()
	if err != nil {
		log.Fatal("error", err)
	}

	msg := LacubaMessage{
		FromName:  "harros",
		Subject:   "New Lacuba!",
		Data:      "This is my message",
		DataMap:   nil,
		LacubaId:  newLacubaId,
		LacubaLat: newLacuba.Latitude,
		LacubaLng: newLacuba.Longtitude,
	}

	token, err := encryptToken(authToken, postFormID)
	if err != nil {
		fmt.Println(err)
	}
	msg.LacubaAuth = token
	app.sendEmail(msg)
}

func (app *Config) getRemoveLacubaForm(c *gin.Context) {
	encrAuth := c.DefaultQuery("auth", "")
	decrypted, err := decryptToken(authToken, encrAuth)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, nil)
	}
	if decrypted == postFormID {
		c.HTML(http.StatusOK, "remove-lacuba-form.gohtml", gin.H{})
	} else {
		c.AbortWithStatusJSON(400, nil)
	}

}

func (app *Config) postRemoveLacubaForm(c *gin.Context) {
	fmt.Println("removing")
}
