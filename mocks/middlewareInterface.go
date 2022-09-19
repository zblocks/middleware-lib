// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	middleware "github.com/sourabhmandal/middleware-lib"
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// middlewareInterface is an autogenerated mock type for the middlewareInterface type
type middlewareInterface struct {
	mock.Mock
}

var MiddlewareMock = &middlewareInterface{}


// GetUserJwtData provides a mock function with given fields: baseUrl, userEmail
func (_m *middlewareInterface) GetUserJwtData(baseUrl string, userEmail string) middleware.GetUserDataByEmailResponse {
	ret := _m.Called(baseUrl, userEmail)

	var r0 middleware.GetUserDataByEmailResponse
	if rf, ok := ret.Get(0).(func(string, string) middleware.GetUserDataByEmailResponse); ok {
		r0 = rf(baseUrl, userEmail)
	} else {
		r0 = ret.Get(0).(middleware.GetUserDataByEmailResponse)
	}

	return r0
}

// SetCors provides a mock function with given fields: _a0
func (_m *middlewareInterface) SetCors(_a0 *gin.Engine) {
	_m.Called(_a0)
}

// VerifyJwtTokenV2 provides a mock function with given fields: c, authServiceBaseUrl
func (_m *middlewareInterface) VerifyJwtTokenV2(c *gin.Context, authServiceBaseUrl string) *middleware.ValidJwt {
	ret := _m.Called(c, authServiceBaseUrl)

	var r0 *middleware.ValidJwt
	if rf, ok := ret.Get(0).(func(*gin.Context, string) *middleware.ValidJwt); ok {
		r0 = rf(c, authServiceBaseUrl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*middleware.ValidJwt)
		}
	}

	return r0
}

type mockConstructorTestingTnewMiddlewareInterface interface {
	mock.TestingT
	Cleanup(func())
}

// newMiddlewareInterface creates a new instance of middlewareInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMiddlewareInterface(t mockConstructorTestingTnewMiddlewareInterface) *middlewareInterface {
	mock := &middlewareInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
