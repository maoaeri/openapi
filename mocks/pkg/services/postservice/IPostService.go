// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	model "github.com/maoaeri/openapi/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// IPostService is an autogenerated mock type for the IPostService type
type IPostService struct {
	mock.Mock
}

// CreatePostService provides a mock function with given fields: post
func (_m *IPostService) CreatePostService(post *model.Post) (int, error) {
	ret := _m.Called(post)

	var r0 int
	if rf, ok := ret.Get(0).(func(*model.Post) int); ok {
		r0 = rf(post)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Post) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAllPostsService provides a mock function with given fields:
func (_m *IPostService) DeleteAllPostsService() (int, error) {
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

// DeletePostService provides a mock function with given fields: postid_param
func (_m *IPostService) DeletePostService(postid_param int) (int, error) {
	ret := _m.Called(postid_param)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(postid_param)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(postid_param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPostsService provides a mock function with given fields: page
func (_m *IPostService) GetAllPostsService(page int) ([]model.Post, int, error) {
	ret := _m.Called(page)

	var r0 []model.Post
	if rf, ok := ret.Get(0).(func(int) []model.Post); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Post)
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

// GetPostService provides a mock function with given fields: postid
func (_m *IPostService) GetPostService(postid int) (*model.Post, int, error) {
	ret := _m.Called(postid)

	var r0 *model.Post
	if rf, ok := ret.Get(0).(func(int) *model.Post); ok {
		r0 = rf(postid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int) int); ok {
		r1 = rf(postid)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(postid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdatePostService provides a mock function with given fields: postid_param, data
func (_m *IPostService) UpdatePostService(postid_param int, data map[string]interface{}) (int, error) {
	ret := _m.Called(postid_param, data)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) int); ok {
		r0 = rf(postid_param, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, map[string]interface{}) error); ok {
		r1 = rf(postid_param, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
