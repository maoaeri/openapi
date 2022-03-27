package usercontrollers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	mocks "github.com/maoaeri/openapi/mocks/pkg/services/userservice"
	jwt_handler "github.com/maoaeri/openapi/pkg"
	"github.com/maoaeri/openapi/pkg/model"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var usertest = &model.User{
	UserID:   1,
	Username: "mao",
	Email:    "mao@user",
	Password: "mao",
}

func TestSignUp(t *testing.T) {
	// create an instance of our test object
	userService := new(mocks.IUserService)

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
			500,
			nil},
		{map[string]interface{}{
			"Username": " ",
			"Email":    "mao1@user",
			"Password": "mao"},
			http.StatusBadRequest,
			errors.New("Username cannot be blank.")},
		{map[string]interface{}{
			"Username": "mao",
			"Email":    "mao1@user",
			"Password": "mao"},
			http.StatusCreated,
			nil},
	}

	for _, test := range userTest {
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
	// create an instance of our test object
	userService := new(mocks.IUserService)

	type Test struct {
		email_param string
		in          map[string]interface{}
		out         int
		err         error
	}
	userTest := []Test{
		{"mao1@user",
			map[string]interface{}{
				"UserID":    1,
				"Usernamea": "mao2",
				"Email":     "mao1@user",
				"Password":  "mao",
				"Role":      "user"},
			500,
			nil},
		{"mao2@user", map[string]interface{}{
			"UserID":   1,
			"Username": "a",
			"Email":    "mao1@user",
			"Password": "mao",
			"Role":     "user"},
			http.StatusBadRequest,
			errors.New("You cannot get other user's information.")},
		{"mao2@user",
			map[string]interface{}{
				"UserID":   1,
				"Username": "mao3",
				"Email":    "mao2@user",
				"Password": "mao",
				"Role":     "user"},
			http.StatusOK,
			nil},
	}
	for _, test := range userTest {
		var testdata *model.User
		mapstructure.Decode(test.in, &testdata)
		authmiddleware := jwt_handler.JwtHandler()
		token, _, _ := authmiddleware.TokenGenerator(testdata)

		//var returnTest *model.User
		//set up expectations
		userService.On("GetUserService", test.email_param, test.in["Email"]).Return(testdata, test.out, test.err)

		userController := UserController{userService}

		// call the code we are testing
		req := httptest.NewRequest("GET", "http://localhost:8080/users/"+test.email_param, nil)
		req.Header = map[string][]string{
			"Authorization": {"Bearer " + token},
		}

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.GET("/users/:email", userController.GetUserHandler)
		engine.ServeHTTP(w, req)

		expectedCode := test.out

		actualCode := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedCode, actualCode)
	}
}

func TestUpdateUser(t *testing.T) {
	// create an instance of our test object
	userService := new(mocks.IUserService)

	type Test struct {
		email_param string
		in          map[string]interface{}
		out         int
	}
	userTest := []Test{
		{"mao1@user",
			map[string]interface{}{
				"Usernamea": "mao2",
				"Email":     "mao1@user",
				"Password":  "mao"},
			500},
		{"mao2",
			map[string]interface{}{
				"Username": "ah",
				"Email":    "mao1",
				"Password": "mao"}, http.StatusBadRequest},
	}

	for _, test := range userTest {
		var testdata *model.User
		mapstructure.Decode(test.in, &testdata)
		authmiddleware := jwt_handler.JwtHandler()
		token, _, _ := authmiddleware.TokenGenerator(testdata)
		//set up expectations
		userService.On("UpdateUserService", test.email_param, test.in["Email"], test.in).Return(test.out, nil)
		userController := UserController{
			userService,
		}

		// call the code we are testing
		b, _ := json.Marshal(test.in)
		body := strings.NewReader(string(b))
		req := httptest.NewRequest("PUT", "http://localhost:8080/users/"+test.email_param, body)
		req.Header = map[string][]string{
			"Authorization": {"Bearer " + token},
		}

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		//path := fmt.Sprintf("/users/%s", test.email_param)
		engine.PUT("/users/:email", userController.UpdateUserHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestDeleteUser(t *testing.T) {
	// create an instance of our test object
	userService := new(mocks.IUserService)

	type Test struct {
		email_param string
		in          map[string]interface{}
		out         int
	}
	userTest := []Test{
		{"mao1@user",
			map[string]interface{}{
				"Username": "mao2",
				"Email":    "mao1@user",
				"Password": "mao",
				"Role":     "user"},
			200},
		{"mao2",
			map[string]interface{}{
				"Username": "ah",
				"Email":    "mao1",
				"Password": "mao",
				"Role":     "user"}, http.StatusBadRequest},
	}

	for _, test := range userTest {
		var testdata *model.User
		mapstructure.Decode(test.in, &testdata)
		authmiddleware := jwt_handler.JwtHandler()
		token, _, _ := authmiddleware.TokenGenerator(testdata)
		//set up expectations
		userService.On("DeleteUserService", test.email_param, test.in["Email"]).Return(test.out, nil)
		userController := UserController{
			userService,
		}

		// call the code we are testing
		req := httptest.NewRequest("DELETE", "http://localhost:8080/users/"+test.email_param, nil)
		req.Header = map[string][]string{
			"Authorization": {"Bearer " + token},
		}

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		//path := fmt.Sprintf("/users/%s", test.email_param)
		engine.DELETE("/users/:email", userController.DeleteUserHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestDeleteAllUsers(t *testing.T) {
	// create an instance of our test object
	userService := new(mocks.IUserService)

	type Test struct {
		in  map[string]interface{}
		out int
	}
	userTest := []Test{
		{
			map[string]interface{}{
				"Username": "mao2",
				"Email":    "mao1@user",
				"Password": "mao",
				"Role":     "admin"},
			200},
	}

	for _, test := range userTest {
		var testdata *model.User
		mapstructure.Decode(test.in, &testdata)
		authmiddleware := jwt_handler.JwtHandler()
		token, _, _ := authmiddleware.TokenGenerator(testdata)
		//set up expectations
		userService.On("DeleteAllUsersService").Return(test.out, nil)
		userController := UserController{
			userService,
		}

		// call the code we are testing
		req := httptest.NewRequest("DELETE", "http://localhost:8080/users", nil)
		req.Header = map[string][]string{
			"Authorization": {"Bearer " + token},
		}

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		//path := fmt.Sprintf("/users/%s", test.email_param)
		engine.DELETE("/users", userController.DeleteAllUsersHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}
