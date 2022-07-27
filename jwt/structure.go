package jwt

import (
	"github.com/gin-gonic/gin"
)

var JWTMiddlewares middlewareInterface = &middlewareStruct{}

type middlewareStruct struct{}
type middlewareInterface interface {
	GetUserID(baseUrl string, userEmail string) UserData
	VerifyJwtTokenV2(c *gin.Context, authServiceBaseUrl string) ValidJwt
}
