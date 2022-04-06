package bearertoken

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler func(*gin.Context)

type Options struct {
	OnAuthorizationHeaderMissing Handler
	OnAuthorizationHeaderInvalid Handler
	OnTokenInvalid               Handler
}

func Middleware(token string, opt ...Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Authorization header is missing in HTTP request
		if authHeader == "" {
			if len(opt) > 0 && opt[0].OnAuthorizationHeaderMissing != nil {
				opt[0].OnAuthorizationHeaderMissing(c)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			return
		}

		authTokens := strings.Split(authHeader, " ")

		// The value of authorization header is invalid
		// It should start with "Bearer ", then the token value
		if len(authTokens) != 2 || authTokens[0] != "Bearer" {
			if len(opt) > 0 && opt[0].OnAuthorizationHeaderInvalid != nil {
				opt[0].OnAuthorizationHeaderInvalid(c)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			return
		}

		// Token value is invalid
		if authTokens[1] != token {
			if len(opt) > 0 && opt[0].OnTokenInvalid != nil {
				opt[0].OnTokenInvalid(c)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			return
		}

		// Everything looks fine, process next action
		c.Next()
	}
}
