package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthentication() gin.HandlerFunc {
	authToken := ">4t4Q&|@Ik8zw;r9%6"
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader != authToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		c.Next()
	}
}
