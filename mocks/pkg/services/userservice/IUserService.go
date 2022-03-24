// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	model "github.com/maoaeri/openapi/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// IUserService is an autogenerated mock type for the IUserService type
type IUserService struct {
	mock.Mock
}

// DeleteAllUsersService provides a mock function with given fields:
func (_m *IUserService) DeleteAllUsersService() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserService provides a mock function with given fields: email_param, email_token
func (_m *IUserService) DeleteUserService(email_param string, email_token string) (int, error) {
	ret := _m.Called(email_param, email_token)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(email_param, email_token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email_param, email_token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsersService provides a mock function with given fields: page
func (_m *IUserService) GetAllUsersService(page int) ([]model.User, int, error) {
	ret := _m.Called(page)

	var r0 []model.User
	if rf, ok := ret.Get(0).(func(int) []model.User); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUserService provides a mock function with given fields: email_param, email_token
func (_m *IUserService) GetUserService(email_param string, email_token string) (*model.User, int, error) {
	ret := _m.Called(email_param, email_token)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string, string) *model.User); ok {
		r0 = rf(email_param, email_token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string, string) int); ok {
		r1 = rf(email_param, email_token)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email_param, email_token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SignUpService provides a mock function with given fields: user
func (_m *IUserService) SignUpService(user *model.User) (int, error) {
	ret := _m.Called(user)

	var r0 int
	if rf, ok := ret.Get(0).(func(*model.User) int); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserService provides a mock function with given fields: email_param, email_token, data
func (_m *IUserService) UpdateUserService(email_param string, email_token string, data map[string]interface{}) (int, error) {
	ret := _m.Called(email_param, email_token, data)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}) int); ok {
		r0 = rf(email_param, email_token, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, map[string]interface{}) error); ok {
		r1 = rf(email_param, email_token, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}