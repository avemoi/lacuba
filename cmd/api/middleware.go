package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const authToken string = ">4t4Q&|@Ik8zw;r9%6"

func TokenAuthentication() gin.HandlerFunc {
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

func QueryParamAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authq := c.Query("auth")
		decrypted, err := decryptToken(authToken, authq)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		if decrypted != postFormID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		c.Next()
	}
}
