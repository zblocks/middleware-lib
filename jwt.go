package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (m *middlewareStruct) GetUserID(baseUrl string) (GetUserIDResponse){
	api := baseUrl + "/getUserPublicProfile"
	response, err := http.Get(api)
	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data GetUserIDResponse 
	json.Unmarshal(responseData, &data)
	return data
}