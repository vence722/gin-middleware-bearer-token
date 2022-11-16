package bearertoken

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenVerificationFunc func(string, *gin.Context) bool

type Handler func(*gin.Context)

type Options struct {
	OnAuthorizationHeaderMissing Handler
	OnAuthorizationHeaderInvalid Handler
	OnTokenInvalid               Handler
}

func Middleware(tokenVerificationFunc TokenVerificationFunc, opt ...Options) gin.HandlerFunc {
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

		// Check token value is valid or not
		if !tokenVerificationFunc(authTokens[1], c) {
			return
		}

		// Everything looks fine, process next action
		c.Next()
	}
}

func MiddlewareWithStaticToken(token string, opt ...Options) gin.HandlerFunc {
	return Middleware(func(s string, c *gin.Context) bool {
		if s != token {
			if len(opt) > 0 && opt[0].OnTokenInvalid != nil {
				opt[0].OnTokenInvalid(c)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			return false
		}
		return true
	}, opt...)
}
