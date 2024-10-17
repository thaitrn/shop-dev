package middlewares

import (
	"shop-dev/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrorCodeTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
