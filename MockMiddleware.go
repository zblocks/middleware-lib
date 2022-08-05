// Code generated by mockery v2.10.0. DO NOT EDIT.

package middleware

import (
	gin "github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"

	mock "github.com/stretchr/testify/mock"
)

// mockMiddleware is an autogenerated mock type for the mockMiddleware type
type mockMiddleware struct {
	mock.Mock
}

var MockMiddleware = &mockMiddleware{}

// SetCors provides a mock function with given fields: _a0
func (_m *mockMiddleware) SetCors(_a0 *gin.Engine) {
	_m.Called(_a0)
}

// VerifyJwtToken provides a mock function with given fields: _a0, _a1
func (_m *mockMiddleware) VerifyJwtToken(_a0 *gin.Context, _a1 string) (bool, jwt.MapClaims, int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*gin.Context, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 jwt.MapClaims
	if rf, ok := ret.Get(1).(func(*gin.Context, string) jwt.MapClaims); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(jwt.MapClaims)
		}
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(*gin.Context, string) int); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*gin.Context, string) error); ok {
		r3 = rf(_a0, _a1)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}
