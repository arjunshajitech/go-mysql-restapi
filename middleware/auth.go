package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

var apiToken = "dummy_token"

func Logger(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.IndentedJSON(400, gin.H{"msg": "Bearer token not found."})
		c.Abort()
		return
	}

	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
		if token == apiToken {
			c.Next()
		} else {
			c.IndentedJSON(400, gin.H{"msg": "Bearer token is invalid."})
			c.Abort()
			return
		}
	} else {
		c.IndentedJSON(400, gin.H{"msg": "Bearer token not found."})
		c.Abort()
		return
	}
}
