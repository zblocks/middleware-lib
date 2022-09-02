package middleware

type GetUserIDResponse struct {
	Status bool `json:"status"`
	UserId int  `json:"userId"`
}

type GetUserDataResponse struct {
	Status bool `json:"status"`
	Data   struct {
		UserIdPk    int    `json:"userId"`
		Email       string `json:"email"`
		Designation string `json:"designation"`
	} `json:"data"`
}

type VerifyJwtTokenResponse struct {
	Status bool     `json:"status"`
	Data   ValidJwt `json:"data"`
}

type ValidJwt struct {
	IsValid bool `json:"isValid"`
}
