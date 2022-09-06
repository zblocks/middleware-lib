package middleware

type GetUserIDResponse struct {
	Status        bool   `json:"status"`
	UserId        int64  `json:"userId"`
	Email         string `json:"email"`
	Designation   string `json:"designation"`
	UserRole      string `json:"userRole"`
	OrgDomainName string `json:"orgDomainName"`
	OrgType       string `json:"orgType"`
}

type GetUserDataResponse struct {
	Status bool `json:"status"`
	Data   struct {
		UserIdPk      int64  `json:"userId"`
		Email         string `json:"email"`
		Designation   string `json:"designation"`
		UserRole      string `json:"userRole"`
		OrgDomainName string `json:"orgDomainName"`
		OrgType       string `json:"orgType"`
	} `json:"data"`
}

type VerifyJwtTokenResponse struct {
	Status bool     `json:"status"`
	Data   ValidJwt `json:"data"`
}

type ValidJwt struct {
	IsValid bool `json:"isValid"`
}
