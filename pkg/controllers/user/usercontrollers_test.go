package usercontrollers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	mocks "github.com/maoaeri/openapi/mocks/pkg/services/userservice"
	"github.com/maoaeri/openapi/pkg/model"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var errBadRequest = errors.New("bad request error")
var errInternalServer = errors.New("internal server error")

func TestSignUp(t *testing.T) {

	type Test struct {
		in  map[string]interface{}
		out int
		err error
	}
	userTest := []Test{
		{map[string]interface{}{
			"Usernamea": "mao2",
			"Email":     "mao1@user",
			"Password":  "mao"},
			http.StatusInternalServerError,
			errInternalServer},
		{map[string]interface{}{
			"Username": " ",
			"Email":    "mao1@user",
			"Password": "mao"},
			http.StatusBadRequest,
			errBadRequest},
		{map[string]interface{}{
			"Username": "mao",
			"Email":    "mao1@user",
			"Password": "mao"},
			http.StatusCreated,
			nil},
	}

	for _, test := range userTest {
		// create an instance of our test object
		userService := new(mocks.IUserService)

		var testdata *model.User
		mapstructure.Decode(test.in, &testdata)
		//set up expectations
		userService.On("SignUpService", testdata).Return(test.out, test.err)
		userController := UserController{
			userService,
		}

		// call the code we are testing
		b, _ := json.Marshal(test.in)
		body := strings.NewReader(string(b))
		req := httptest.NewRequest("POST", "http://localhost:8080/users/signup", body)
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.POST("/users/signup", userController.SignUpHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestGetUser(t *testing.T) {

	type Test struct {
		in  map[string]interface{}
		out int
		err error
	}
	userTest := []Test{
		{map[string]interface{}{
			"UserID":    1,
			"Usernamea": "mao2",
			"Email":     "mao1@user",
			"Password":  "mao",
			"Role":      "user"},
			http.StatusInternalServerError,
			errInternalServer},
		{map[string]interface{}{
			"UserID":   1,
			"Username": "a",
			"Email":    "mao1@user",
			"Password": "mao",
			"Role":     "user"},
			http.StatusBadRequest,
			errBadRequest},
		{map[string]interface{}{
			"UserID":   1,
			"Username": "mao3",
			"Email":    "mao2@user",
			"Password": "mao",
			"Role":     "user"},
			http.StatusOK,
			nil},
	}

	var userid = 1
	for _, test := range userTest {
		// create an instance of our test object
		userService := new(mocks.IUserService)

		var testdata *model.User
		mapstructure.Decode(test.in, &testdata)

		//var returnTest *model.User
		//set up expectations
		userService.On("GetUserService", userid).Return(testdata, test.out, test.err)

		userController := UserController{userService}

		// call the code we are testing
		req := httptest.NewRequest("GET", "http://localhost:8080/users/1", nil)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.GET("/users/:userid", userController.GetUserHandler)
		engine.ServeHTTP(w, req)

		expectedCode := test.out

		actualCode := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedCode, actualCode)
	}
}

func TestGetAllUsers(t *testing.T) {

	type Test struct {
		page_param string
		//in         map[string]interface{}
		out int
		err error
	}
	userTest := []Test{
		{"1", http.StatusOK, nil},
		{"1", http.StatusInternalServerError, errInternalServer},
	}

	var returnTest []model.User

	for _, test := range userTest {
		// create an instance of our test object
		userService := new(mocks.IUserService)

		a, _ := strconv.Atoi(test.page_param)

		//set up expectations
		userService.On("GetAllUsersService", a).Return(returnTest, test.out, test.err)

		userController := UserController{userService}

		// call the code we are testing
		req := httptest.NewRequest("GET", "http://localhost:8080/users?page="+test.page_param, nil)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.GET("/users", userController.GetAllUsersHandler)
		engine.ServeHTTP(w, req)

		expectedCode := test.out

		actualCode := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedCode, actualCode)
	}
}

func TestUpdateUser(t *testing.T) {

	type Test struct {
		in  map[string]interface{}
		out int
		err error
	}
	userTest := []Test{
		{map[string]interface{}{
			"Email": "mao1@user"},
			http.StatusInternalServerError,
			errInternalServer},
		{map[string]interface{}{
			"Emaila": "maoo2@user"},
			http.StatusBadRequest,
			errBadRequest},
		{map[string]interface{}{
			"Email": "mao3@user"},
			http.StatusOK,
			nil},
	}

	for _, test := range userTest {
		// create an instance of our test object
		userService := new(mocks.IUserService)

		//set up expectations
		userService.On("UpdateUserService", 0, test.in).Return(test.out, test.err)
		userController := UserController{
			userService,
		}

		// call the code we are testing
		b, _ := json.Marshal(test.in)
		body := strings.NewReader(string(b))
		req := httptest.NewRequest("PUT", "http://localhost:8080/users/0", body)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.PUT("/users/:userid", userController.UpdateUserHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestDeleteUser(t *testing.T) {

	type Test struct {
		out int
		err error
	}
	userTest := []Test{
		{http.StatusOK, nil},
		{http.StatusInternalServerError, errInternalServer},
	}

	var userid = 0
	for _, test := range userTest {
		// create an instance of our test object
		userService := new(mocks.IUserService)

		//set up expectations
		userService.On("DeleteUserService", userid).Return(test.out, test.err)
		userController := UserController{
			userService,
		}

		// call the code we are testing
		req := httptest.NewRequest("DELETE", "http://localhost:8080/users/0", nil)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		//path := fmt.Sprintf("/users/%s", test.email_param)
		engine.DELETE("/users/:userid", userController.DeleteUserHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestDeleteAllUsers(t *testing.T) {

	type Test struct {
		out int
		err error
	}
	userTest := []Test{
		{http.StatusOK, nil},
		{http.StatusInternalServerError, errInternalServer},
	}

	for _, test := range userTest {
		// create an instance of our test object
		userService := new(mocks.IUserService)

		//set up expectations
		userService.On("DeleteAllUsersService").Return(test.out, test.err)
		userController := UserController{
			userService,
		}

		// call the code we are testing
		req := httptest.NewRequest("DELETE", "http://localhost:8080/users", nil)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.DELETE("/users", userController.DeleteAllUsersHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}
