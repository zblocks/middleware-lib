package jwt

import (
	"github.com/gin-gonic/gin"
)

var JWTMiddlewares middlewareInterface = &middlewareStruct{}

type middlewareStruct struct{}
type middlewareInterface interface {
	GetUserID(userServiceBaseUrl string, userEmail string) UserData
	VerifyJwtTokenV2(c *gin.Context, authServiceBaseUrl string) bool
}
