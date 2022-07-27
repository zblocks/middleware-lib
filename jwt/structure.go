package jwt

var JWTMiddlewares middlewareInterface = &middlewareStruct{}

type middlewareStruct struct{}
type middlewareInterface interface {
	GetUserID(baseUrl string, userEmail string) UserData
}
