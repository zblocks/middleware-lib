package middleware

import (
	"errors"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyJwtToken(c *gin.Context, jwtSecret string) (bool, jwt.MapClaims, int, error) {
	auth_token := c.Request.Header["Authorization"]

	// no auth token error
	if len(auth_token) == 0 {
		return false, nil, ErrAuthorizationTokenEmpty, errors.New(AuthorizationTokenEmpty)
	}

	jwttoken := strings.Split(auth_token[0], " ")[1]

	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(jwttoken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(AuthorizationTokenInvalid)
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return false, nil, ErrAuthorizationTokenInvalid, errors.New(AuthorizationTokenInvalid)
	}
	if !tkn.Valid {
		return false, nil, ErrAuthorizationTokenInvalid, errors.New(AuthorizationTokenInvalid)
	}
	return true, claims, 0, nil
}

func SetCors(url string) {
	// set cors access
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{url},
		AllowMethods: []string{"PUT", "PATCH, GET, POST, DELETE, OPTIONS"},
		AllowHeaders: []string{"*"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.Run()
}
