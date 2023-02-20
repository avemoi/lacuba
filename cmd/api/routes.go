package main

import "github.com/gin-gonic/gin"

func (app *Config) GetRoutes() *gin.Engine {
	router := gin.Default()
	lacuba := router.Group("/lacuba", TokenAuthentication())
	{
		lacuba.GET("/list", app.getLacubas)
		lacuba.POST("/", app.addLacuba)
	}

	return router
}
