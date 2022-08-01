package middleware

type UserData struct {
	UserIdPk    int    `json:"userId"`
	Email       string `json:"email"`
	Designation string `json:"designation"`
}

type GetUserIDResponse struct {
	Status bool     `json:"status"`
	Data   UserData `json:"data"`
}

type VerifyJwtTokenResponse struct {
	Status bool     `json:"status"`
	Data   ValidJwt `json:"data"`
}

type ValidJwt struct {
	IsValid bool `json:"isValid"`
}