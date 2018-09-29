package middleware

import (
	"github.com/gin-gonic/gin"
	"template/handler"
	"template/pkg/errno"
	"template/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		t, err := token.ParseRequest(c)
		if err != nil {

			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set("token", t)
		c.Next()
	}
}
