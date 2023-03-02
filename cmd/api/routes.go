package main

import "github.com/gin-gonic/gin"

func (app *Config) GetRoutes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("./cmd/templates/*")
	lacuba := router.Group("/lacuba", TokenAuthentication())
	{
		lacuba.GET("/list", app.getLacubas)
		lacuba.POST("/", app.addLacuba)
	}
	actions := router.Group("/lacuba")
	{
		actions.GET("/remove-lacuba", app.getRemoveLacubaForm)
		actions.POST("/remove-lacuba", app.postRemoveLacubaForm)
	}
	return router
}
