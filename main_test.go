package middleware_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	middleware "github.com/sourabhmandal/test-zbyte-middleware"
	middlewareMock "github.com/sourabhmandal/test-zbyte-middleware/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMockFunctions(t *testing.T) {
	middleware.MiddlewareHandler = middlewareMock.MiddlewareMock
	jwtToken := "test_jwt_string"
	ctx := &gin.Context{}
	eng := &gin.Engine{}
	auth_base_url := "auth_base_url"
	test_email := "test@gmail.com"
	middlewareMock.MiddlewareMock.On("VerifyJwtTokenV2", ctx, jwtToken).Return(true)
	middlewareMock.MiddlewareMock.On("SetCors", eng).Return(true)
	middlewareMock.MiddlewareMock.On("GetUserID", auth_base_url, test_email).Return(middleware.GetUserIDResponse{
		Status: true,
		UserId: 80,
	})

	t.Run("When verify token mock function given correct data", func(t *testing.T) {
		resp := middlewareMock.MiddlewareMock.VerifyJwtTokenV2(ctx, jwtToken)
		assert.Equal(t, true, resp)
	})

	t.Run("When SetCors() function given correct data", func(t *testing.T) {
		middlewareMock.MiddlewareMock.SetCors(eng)
	})

	t.Run("When GetUserID() function given correct data", func(t *testing.T) {
		resp := middlewareMock.MiddlewareMock.GetUserID(auth_base_url, test_email)
		assert.Equal(t, true, resp.Status)
		assert.Equal(t, 80, resp.UserId)
	})
}
