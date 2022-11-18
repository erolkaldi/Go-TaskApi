package middleware

import (
	"net/http"
	"strings"
	"task-api/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Abort()
			return
		}
		array := strings.Split(token, " ")
		token = array[1]
		claims, msg := helpers.ValidateToken(token)
		if msg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			c.Abort()
		}
		c.Set("email", claims.Email)
		c.Set("username", claims.UserName)
		c.Set("userid", claims.UserID)
		c.Set("usertype", claims.UserType)
		c.Next()
	}
}
