# gin-middleware-bearer-token
Middleware of bearer token authentication for Gin (one of the best Go HTTP Web framework).

## Installation
```shell
go get github.com/vence722/gin-middleware-bearer-token
```

## Usage

### Static token
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/vence722/gin-middleware-bearer-token"
)

var AuthToken = "auth_token"

func main() {
    r := gin.Default()
	
    r.Use(bearertoken.MiddlewareWithStaticToken(AuthToken))
	
    r.GET("/", func(c *gin.Context) {
	    c.String(200, "ok")
    })
	
    r.Run(":8080")
}
```

### Token verification
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/vence722/gin-middleware-bearer-token"
)

func isTokenValid(token string) bool {
	// Custom token verification logic here
	// Check token in DB/config file
	// If token is not valid, return false
	return true
}

func main() {
    r := gin.Default()
	
    r.Use(bearertoken.Middleware(func (token string, c *gin.Context) bool {
		if !isTokenValid(token) {
			// If token is not valid, set HTTP status code to 401 (Unauthorized)
			// And return false for blocking the request flow
			c.AbortWithStatus(http.StatusUnauthorized)
			return false
		}
		// If token is valid, return true to continue the request flow
		return true
	}))
	
    r.GET("/", func(c *gin.Context) {
	    c.String(200, "ok")
    })
	
    r.Run(":8080")
}
```

### Customize handlers with static token
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/vence722/gin-middleware-bearer-token"
)

var AuthToken = "auth_token"

func main() {
    r := gin.Default()
	
    r.Use(bearertoken.MiddlewareWithStaticToken(AuthToken, bearertoken.Options{
		OnTokenInvalid: func(c *gin.Context) {
			// Triggered when auth token doesn't 
			// match the value in Authorization header
			c.AbortWithStatus(401)
		},
		OnAuthorizationHeaderInvalid: func(c *gin.Context) {
			// Triggered when Authorization Header
			// is not in "Bearer xxx" format
			c.AbortWithStatus(400)
		},
		OnAuthorizationHeaderMissing: func(c *gin.Context) {
			// Triggered when Authorization Header
			// is missing in HTTP request
			c.AbortWithStatus(400)
		},
	}))
	
    r.GET("/", func(c *gin.Context) {
	    c.String(200, "ok")
    })
	
    r.Run(":8080")
}
```
