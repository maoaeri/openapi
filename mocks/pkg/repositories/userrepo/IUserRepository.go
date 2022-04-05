// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	model "github.com/maoaeri/openapi/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// CheckDuplicateEmail provides a mock function with given fields: email
func (_m *IUserRepository) CheckDuplicateEmail(email string) bool {
	ret := _m.Called(email)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreateUser provides a mock function with given fields: user
func (_m *IUserRepository) CreateUser(user *model.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllUsers provides a mock function with given fields:
func (_m *IUserRepository) DeleteAllUsers() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: userid
func (_m *IUserRepository) DeleteUser(userid int) error {
	ret := _m.Called(userid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsers provides a mock function with given fields: page
func (_m *IUserRepository) GetAllUsers(page int) ([]model.User, error) {
	ret := _m.Called(page)

	var r0 []model.User
	if rf, ok := ret.Get(0).(func(int) []model.User); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: userid
func (_m *IUserRepository) GetUser(userid int) (*model.User, error) {
	ret := _m.Called(userid)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(int) *model.User); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: userid, data
func (_m *IUserRepository) UpdateUser(userid int, data map[string]interface{}) error {
	ret := _m.Called(userid, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) error); ok {
		r0 = rf(userid, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
