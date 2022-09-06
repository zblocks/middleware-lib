package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	GetUserJwtData(baseUrl string, userEmail string) GetUserIDResponse
	VerifyJwtTokenV2(c *gin.Context, authServiceBaseUrl string) bool
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

func (m *middlewareStruct) GetUserJwtData(userServiceBaseUrl string, userEmail string) GetUserIDResponse {
	api, _ := url.JoinPath(userServiceBaseUrl, "/getUserProfileByEmail")
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

	var data GetUserDataResponse
	derr := json.NewDecoder(res.Body).Decode(&data)
	if derr != nil {
		panic(derr.Error())
	}
	resp := GetUserIDResponse{
		Status: data.Status,
		UserId: data.Data.UserIdPk,
		Email: data.Data.Email,
		Designation: data.Data.Designation,
		UserRole: data.Data.UserRole,
		OrgDomainName: data.Data.OrgDomainName,
		OrgType: data.Data.OrgType,
	}
	if data.Status {
		return resp
	}
	return GetUserIDResponse{}
}

func (m *middlewareStruct) VerifyJwtTokenV2(c *gin.Context, authServiceBaseUrl string) bool {
	api, _ := url.JoinPath(authServiceBaseUrl, "/auth/verifyToken")
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
