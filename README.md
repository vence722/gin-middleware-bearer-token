# gin-middleware-bearer-token
Middleware of bearer token authentication for Gin (one of the best Go HTTP Web framework).

## Installation
```shell
go get github.com/vence722/gin-middleware-bearer-token
```

## Usage

### Basic
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/vence722/gin-middleware-bearer-token"
)

var AuthToken = "auth_token"

func main() {
    r := gin.Default()
	
    r.Use(bearertoken.Middleware(AuthToken))
	
    r.GET("/", func(c *gin.Context) {
	    c.String(200, "ok")
    })
	
    r.Run(":8080")
}
```

### Customize handlers
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/vence722/gin-middleware-bearer-token"
)

var AuthToken = "auth_token"

func main() {
    r := gin.Default()
	
    r.Use(bearertoken.Middleware(AuthToken, bearertoken.Options{
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
