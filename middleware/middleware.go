package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var MiddlewareHandler middlewareInterface = &middlewareStruct{}

type middlewareStruct struct{}

type middlewareInterface interface {
	SetCors(*gin.Engine)
	GetUserJwtData(baseUrl string, userEmail string) GetUserDataByEmailResponse
	UseDepricatedVerifyJwtToken(c *gin.Context, authServiceBaseUrl string) *ValidJwt
	KeycloakTokenVerify(accessToken string, keycloakBaseUrl string) (jwtResponse *VerifyJwtTokenResponseKeycloak, errorData error) 
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

func (m *middlewareStruct) GetUserJwtData(userServiceBaseUrl string, userEmail string) GetUserDataByEmailResponse {
	api, _ := url.JoinPath(userServiceBaseUrl, "/getUserProfileByEmail")
	log.Printf(`USER SERVICE BASE URL :: %s / USER EMAIL :: %s / CALLING API :: %s`, userServiceBaseUrl, userEmail, api)
	body := []byte(fmt.Sprintf(`{"userEmail": "%s"}`, userEmail))

	r, err := http.NewRequest("POST", api, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("CREATE API REQUEST ERROR :: %v", err)
		return GetUserDataByEmailResponse{
			Status: false,
		}
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	log.Printf("CALLED API :: %s", api)
	if err != nil {
		log.Printf("RECIEVED API ERROR :: %v", err)
	}

	defer res.Body.Close()

	var data GetUserDataByEmailResponse
	log.Println("parsing json response")
	derr := json.NewDecoder(res.Body).Decode(&data)
	if derr != nil {
		log.Printf("JSON PARSER ERROR :: %v", derr)
		return GetUserDataByEmailResponse{
			Status: false,
		}
	}

	if data.Status {
		return data
	}
	return GetUserDataByEmailResponse{
		Status: false,
	}
}

func (m *middlewareStruct) UseDepricatedVerifyJwtToken(c *gin.Context, authServiceBaseUrl string) *ValidJwt {
	api, _ := url.JoinPath(authServiceBaseUrl, "/v1/auth/verifyToken")
	log.Printf(`AUTH SERVICE BASE URL :: %s / CALLING API %s`, authServiceBaseUrl, api)
	r, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Printf("CREATE API REQUEST ERROR :: %v", err)
		return &ValidJwt{
			IsValid: false,
		}
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", c.Request.Header["Authorization"][0])

	client := &http.Client{}
	res, err := client.Do(r)
	log.Printf("CALLED API :: %s", api)
	if err != nil {
		log.Printf("RECIEVED API ERROR :: %v", err)
		return &ValidJwt{
			IsValid: false,
		}
	}

	defer res.Body.Close()

	var data VerifyJwtTokenResponse
	log.Println("parsing json response")
	derr := json.NewDecoder(res.Body).Decode(&data)
	if derr != nil {
		log.Printf("JSON PARSER ERROR :: %v", derr)
		return &ValidJwt{
			IsValid: false,
		}
	}
	if data.Status {
		return data.Data
	}
	return &ValidJwt{
		IsValid: false,
	}
}
