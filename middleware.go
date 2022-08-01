package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var MiddlewareHandler middlewareInterface = &middlewareStruct{}

type middlewareStruct struct{}

type middlewareInterface interface {
	VerifyJwtToken(*gin.Context, string) (bool, jwt.MapClaims, int, error)
	SetCors(*gin.Engine)
	GetUserID(baseUrl string, userEmail string) UserData
	VerifyJwtTokenV2(c *gin.Context, authServiceBaseUrl string) bool
}

func (m *middlewareStruct) VerifyJwtToken(c *gin.Context, jwtSecret string) (bool, jwt.MapClaims, int, error) {
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

func (m *middlewareStruct) SetCors(r *gin.Engine) {
	// set cors access
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))
}



func (m *middlewareStruct) GetUserID(baseUrl string, userEmail string) UserData {
	api := baseUrl + "/getUserProfileByEmail"
	body := []byte(fmt.Sprintf(`{"userEmail": "%s"}`, userEmail))

	r, err := http.NewRequest("POST", api, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var data GetUserIDResponse
	derr := json.NewDecoder(res.Body).Decode(&data)
	if derr != nil {
		panic(derr.Error())
	}
	if data.Status {
		return data.Data
	}
	return UserData{}
}


func (m *middlewareStruct) VerifyJwtTokenV2(c *gin.Context, authServiceBaseUrl string) bool {
	api := authServiceBaseUrl + "/auth/verifyToken"
	r, err := http.NewRequest("GET", api, nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", c.Request.Header["Authorization"][0])

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var data VerifyJwtTokenResponse
	derr := json.NewDecoder(res.Body).Decode(&data)
	if derr != nil {
		panic(derr.Error())
	}
	if data.Status {
		return data.Data.IsValid
	}
	return false
}