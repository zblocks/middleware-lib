package middleware

import (
	gin "github.com/gin-gonic/gin"
)

var MiddlewareHandler middlewareInterface = &middlewareStruct{}

type middlewareStruct struct{}
type middlewareInterface interface {
	GetUserID(baseUrl string) GetUserIDResponse
	SetCors(r *gin.Engine)
}
