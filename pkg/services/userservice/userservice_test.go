package userservice

import (
	"errors"
	"net/http"
	"testing"

	mocks "github.com/maoaeri/openapi/mocks/pkg/repositories/userrepo"
	"github.com/maoaeri/openapi/pkg/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var errBadRequest = errors.New("bad request error")
var errInternalServer = errors.New("internal server error")

func TestSignUp(t *testing.T) {
	type Test struct {
		in               *model.User
		checkEmailResult bool
		out              int
		err              error
	}

	var userTest = []Test{
		{
			&model.User{
				Username: "",
				Email:    "a",
				Password: "a",
			}, false, http.StatusBadRequest, errBadRequest,
		},
		{
			&model.User{
				Username: "a",
				Email:    "",
				Password: "a",
			}, false, http.StatusBadRequest, errBadRequest,
		},
		{
			&model.User{
				Username: "a",
				Email:    "a",
				Password: "",
			}, false, http.StatusBadRequest, errBadRequest,
		},
		{
			&model.User{
				Username: "a",
				Email:    "a",
				Password: "a",
			}, true, http.StatusBadRequest, errBadRequest,
		},
		{
			&model.User{
				Username: "a",
				Email:    "a",
				Password: "a",
			}, false, http.StatusInternalServerError, errInternalServer,
		},
		{
			&model.User{
				Username: "a",
				Email:    "a",
				Password: "a",
			}, false, http.StatusCreated, nil,
		},
	}

	for _, test := range userTest {
		// create an instance of our test object
		userRepo := new(mocks.IUserRepository)

		//set up expectations
		userRepo.On("CheckDuplicateEmail", test.in.Email).Return(test.checkEmailResult)
		userRepo.On("CreateUser", test.in).Return(test.err)
		userService := UserService{
			userRepo,
		}

		got, _ := userService.SignUpService(test.in)
		assert.Equal(t, test.out, got)
	}
}

func TestGetAllUsers(t *testing.T) {
	type Test struct {
		out int
		err error
	}

	userTest := []Test{
		{http.StatusBadRequest, gorm.ErrRecordNotFound},
		{http.StatusInternalServerError, errInternalServer},
		{http.StatusOK, nil},
	}

	var usersData = []model.User{}
	for _, test := range userTest {
		// create an instance of our test object
		userRepo := new(mocks.IUserRepository)

		//set up expectations
		userRepo.On("GetAllUsers", 0).Return(usersData, test.err) //page = 0 for all tests
		userService := UserService{
			userRepo,
		}

		_, got, _ := userService.GetAllUsersService(0)
		assert.Equal(t, test.out, got)
	}
}

func TestGetUser(t *testing.T) {
	type Test struct {
		out int
		err error
	}

	userTest := []Test{
		{http.StatusBadRequest, gorm.ErrRecordNotFound},
		{http.StatusInternalServerError, errInternalServer},
		{http.StatusOK, nil},
	}

	var userid = 0
	var userData = &model.User{}
	for _, test := range userTest {
		// create an instance of our test object
		userRepo := new(mocks.IUserRepository)

		//set up expectations
		userRepo.On("GetUser", userid).Return(userData, test.err)
		userService := UserService{
			userRepo,
		}

		_, got, _ := userService.GetUserService(userid)
		assert.Equal(t, test.out, got)
	}
}

func TestUpdateUser(t *testing.T) {
	type Test struct {
		data map[string]interface{}
		out  int
		err  error
	}

	userTest := []Test{
		{map[string]interface{}{"userid": 1}, http.StatusBadRequest, errBadRequest},
		{map[string]interface{}{"userid": 0}, http.StatusInternalServerError, errInternalServer},
		{map[string]interface{}{"userid": 0}, http.StatusOK, nil},
	}

	var userid = 0
	for _, test := range userTest {
		// create an instance of our test object
		userRepo := new(mocks.IUserRepository)

		//set up expectations
		userRepo.On("UpdateUser", userid, test.data).Return(test.err)
		userService := UserService{
			userRepo,
		}

		got, _ := userService.UpdateUserService(userid, test.data)
		assert.Equal(t, test.out, got)
	}
}

func TestDeleteAllUsers(t *testing.T) {
	type Test struct {
		out int
		err error
	}

	userTest := []Test{
		{http.StatusInternalServerError, errInternalServer},
		{http.StatusOK, nil},
	}

	for _, test := range userTest {
		// create an instance of our test object
		userRepo := new(mocks.IUserRepository)

		//set up expectations
		userRepo.On("DeleteAllUsers").Return(test.err)
		userService := UserService{
			userRepo,
		}

		got, _ := userService.DeleteAllUsersService()
		assert.Equal(t, test.out, got)
	}
}

func TestDeleteUser(t *testing.T) {
	type Test struct {
		out int
		err error
	}

	userTest := []Test{
		{http.StatusInternalServerError, errInternalServer},
		{http.StatusOK, nil},
	}

	var userid = 0
	for _, test := range userTest {
		// create an instance of our test object
		userRepo := new(mocks.IUserRepository)

		//set up expectations
		userRepo.On("DeleteUser", userid).Return(test.err)
		userService := UserService{
			userRepo,
		}

		got, _ := userService.DeleteUserService(userid)
		assert.Equal(t, test.out, got)
	}
}
