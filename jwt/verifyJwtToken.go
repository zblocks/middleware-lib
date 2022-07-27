package jwt

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middlewareStruct) VerifyJwtTokenV2(c *gin.Context, authServiceBaseUrl string) ValidJwt {
	api := authServiceBaseUrl + "/auth/refreshToken"
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
		return data.Data
	}
	return ValidJwt{}
}