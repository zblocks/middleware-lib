package jwt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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


