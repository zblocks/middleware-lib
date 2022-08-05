package middleware

type GetUserIDResponse struct {
	Status bool `json:"status"`
	UserId int  `json:"userId"`
}

type VerifyJwtTokenResponse struct {
	Status bool     `json:"status"`
	Data   ValidJwt `json:"data"`
}

type ValidJwt struct {
	IsValid bool `json:"isValid"`
}
