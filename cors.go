package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (m *middlewareStruct) SetCors(r *gin.Engine) {
	// set cors access
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")
	// Register the middleware
	r.Use(cors.New(corsConfig))
}
