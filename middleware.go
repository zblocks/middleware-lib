package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var MiddlewareHandler middlewareInterface = &middlewareStruct{}

type middlewareStruct struct{}

type middlewareInterface interface {
	SetCors(*gin.Engine)
	GetUserID(baseUrl string, userEmail string) GetUserIDResponse
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

func (m *middlewareStruct) GetUserID(baseUrl string, userEmail string) GetUserIDResponse {
	api, _ := url.JoinPath(baseUrl ,"/getUserProfileByEmail")
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
