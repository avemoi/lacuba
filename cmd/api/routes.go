package main

import "github.com/gin-gonic/gin"

func (r *repo) GetRoutes() *gin.Engine {
	router := gin.Default()
	lacuba := router.Group("/lacuba", TokenAuthentication())
	{
		lacuba.GET("/list", r.getLacubas)
		lacuba.POST("/", r.addLacuba)
	}

	return router
}
