## Middleware library to handle jwt

### Work with library

1. import the package using the commands in your project

```go
import (
	"github.com/Zbyteio/middleware-lib"
)

```
Inside terminal run 
```go
go get github.com/Zbyteio/middleware-lib

```
2. Functions included in the library

```go
func VerifyJwtToken(c *gin.Context, jwtSecret string) (bool, jwt.MapClaims, int, error)
```

```go
func SetCors(r *gin.Engine) 
```

3. Variables defined in the library

```go
var ErrAuthorizationTokenEmpty = 20003
var ErrAuthorizationTokenInvalid = 20004

var AuthorizationTokenEmpty = "authorization token not provided"
var AuthorizationTokenInvalid = "authorization token invalid"
```



## Mock middleware in test

### Example mocking middleware functions VerifyJwtTokenV2() in your service
1. In your test function replace MiddlewareHandler interface with the mock interface

```go
middleware.MiddlewareHandler = middlewareMock.MiddlewareMock
```
2. Define mocking behavior for the function

```go
	middlewareMock.MiddlewareMock.On("VerifyJwtTokenV2", mock.AnythingOfType("*gin.Context"), "<auth_service_base_url>").Return(true)

```

3. Use the new keycloak access token verification as follows

```go

srv.GET("/a", func(ctx *gin.Context) {
  access_token := ctx.GetHeader("Authorization")
	isValid, err := middleware.MiddlewareHandler.KeycloakTokenVerify(access_token, "https://appdev.zbyte.io/keycloak-poc/")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, isValid)
})
```