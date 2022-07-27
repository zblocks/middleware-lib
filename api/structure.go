package middleware

import (
	gin "github.com/gin-gonic/gin"
)

var ApiMiddlewares middlewareInterface = &middlewareStruct{}

type middlewareStruct struct{}
type middlewareInterface interface {
	SetCors(r *gin.Engine)
}
