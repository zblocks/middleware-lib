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



