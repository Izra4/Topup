package middleware

import (
	"TopUpWEb/entity"
	"TopUpWEb/sdk"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorize := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(authorize, "Bearer ") {
			c.Abort()
			msg := "wrong header value"
			sdk.FailOrError(c, http.StatusForbidden, msg, errors.New(msg))
			return
		}
		tokenStr := authorize[7:]
		claims := entity.AdminClaims{}
		key := os.Getenv("SECRET_KEY")

		if _, err := sdk.DecodeToken(tokenStr, &claims, key); err != nil {
			c.Abort()
			sdk.FailOrError(c, http.StatusUnauthorized, "unauthorized", err)
			return
		}
		c.Set("admin", claims)
	}
}
