package middleware

type GetUserDataByEmailResponse struct {
	Status bool `json:"status"`
	Data   struct {
		UserId        int64  `json:"userId"`
		Email         string `json:"email"`
		Designation   string `json:"designation"`
		OrgDomainName string `json:"orgDomainName"`
		UserRole      string `json:"userRole"`
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
