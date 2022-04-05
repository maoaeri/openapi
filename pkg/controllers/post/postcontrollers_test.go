package postcontrollers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	mocks "github.com/maoaeri/openapi/mocks/pkg/services/postservice"
	jwt_handler "github.com/maoaeri/openapi/pkg"
	"github.com/maoaeri/openapi/pkg/model"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var errBadRequest = errors.New("bad request error")
var errInternalServer = errors.New("internal server error")

func TestCreatePost(t *testing.T) {

	type Test struct {
		inPost map[string]interface{}
		inUser map[string]interface{}
		out    int
		err    error
	}
	postTest := []Test{
		{map[string]interface{}{
			"contenta": "hihi"},
			map[string]interface{}{
				"UserID": 0},
			http.StatusInternalServerError,
			errInternalServer},
		{map[string]interface{}{
			"content": ""},
			map[string]interface{}{
				"UserID": 0},
			http.StatusBadRequest,
			errBadRequest},
		{map[string]interface{}{
			"content": "hihi"},
			map[string]interface{}{
				"UserID": 0},
			http.StatusCreated,
			nil},
	}

	for _, test := range postTest {
		// create an instance of our test object
		postService := new(mocks.IPostService)

		var testdata *model.Post
		mapstructure.Decode(test.inPost, &testdata)
		//set up expectations
		postService.On("CreatePostService", testdata).Return(test.out, test.err)
		postController := PostController{
			postService,
		}

		var testdata2 *model.User
		mapstructure.Decode(test.inUser, &testdata2)
		authmiddleware := jwt_handler.JwtHandler()
		token, _, _ := authmiddleware.TokenGenerator(testdata2)

		// call the code we are testing
		b, _ := json.Marshal(test.inPost)
		body := strings.NewReader(string(b))
		req := httptest.NewRequest("POST", "http://localhost:8080/users/signup", body)
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.POST("/users/signup", postController.CreatePostHandler)
		req.Header = map[string][]string{
			"Authorization": {"Bearer " + token},
		}

		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestUpdatePost(t *testing.T) {

	type Test struct {
		inPost map[string]interface{}
		out    int
		err    error
	}
	postTest := []Test{
		{map[string]interface{}{
			"contenta": "hihi"},
			http.StatusInternalServerError,
			errInternalServer},
		{map[string]interface{}{
			"content": ""},
			http.StatusBadRequest,
			errBadRequest},
		{map[string]interface{}{
			"content": "hihi"},
			http.StatusOK,
			nil},
	}
	var postid = 0

	for _, test := range postTest {
		// create an instance of our test object
		postService := new(mocks.IPostService)

		//set up expectations
		postService.On("UpdatePostService", postid, test.inPost).Return(test.out, test.err)
		postController := PostController{
			postService,
		}

		// call the code we are testing
		b, _ := json.Marshal(test.inPost)
		body := strings.NewReader(string(b))
		req := httptest.NewRequest("PUT", "http://localhost:8080/posts/0", body)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.PUT("/posts/:postid", postController.UpdatePostHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestDeletePost(t *testing.T) {

	//no need for post info

	type Test struct {
		out int
		err error
	}
	postTest := []Test{
		{http.StatusInternalServerError, errInternalServer},
		{http.StatusOK, nil},
	}

	var postid = 0
	for _, test := range postTest {
		// create an instance of our test object
		postService := new(mocks.IPostService)

		//set up expectations
		postService.On("DeletePostService", postid).Return(test.out, test.err)
		postController := PostController{
			postService,
		}

		// call the code we are testing
		req := httptest.NewRequest("DELETE", "http://localhost:8080/posts/0", nil)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.DELETE("/posts/:postid", postController.DeletePostHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestDeleteAllPosts(t *testing.T) {

	//no need for post info

	type Test struct {
		out int
		err error
	}
	postTest := []Test{
		{http.StatusInternalServerError,
			errInternalServer},
		{http.StatusOK,
			nil},
	}

	for _, test := range postTest {
		// create an instance of our test object
		postService := new(mocks.IPostService)

		//set up expectations
		postService.On("DeleteAllPostsService").Return(test.out, test.err)
		postController := PostController{
			postService,
		}

		// call the code we are testing
		req := httptest.NewRequest("DELETE", "http://localhost:8080/posts", nil)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.DELETE("/posts", postController.DeleteAllPostsHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestGetPost(t *testing.T) {

	//no need for post info

	type Test struct {
		postid_param string
		out          int
		err          error
	}
	postTest := []Test{
		{"1",
			http.StatusBadRequest,
			errBadRequest},
		{"1",
			http.StatusOK,
			nil},
	}

	for _, test := range postTest {
		// create an instance of our test object
		postService := new(mocks.IPostService)

		var returnData *model.Post

		a, _ := strconv.Atoi(test.postid_param)
		//set up expectations
		postService.On("GetPostService", a).Return(returnData, test.out, test.err)
		postController := PostController{
			postService,
		}

		// call the code we are testing
		req := httptest.NewRequest("GET", "http://localhost:8080/posts/"+test.postid_param, nil)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.GET("/posts/:postid", postController.GetPostHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}

func TestGetAllPosts(t *testing.T) {

	//no need for post info

	type Test struct {
		page_param string
		out        int
		err        error
	}
	postTest := []Test{
		{"1",
			http.StatusInternalServerError,
			errInternalServer},
		{"1",
			http.StatusOK,
			nil},
	}

	for _, test := range postTest {
		// create an instance of our test object
		postService := new(mocks.IPostService)

		//	a, _ := strconv.Atoi(test.page_param)
		var returnData []model.Post

		//set up expectations
		//why page param must be 0?
		postService.On("GetAllPostsService", 0).Return(returnData, test.out, test.err)
		postController := PostController{
			postService,
		}

		// call the code we are testing
		req := httptest.NewRequest("GET", "http://localhost:8080/posts", nil)

		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		engine.GET("/posts", postController.GetAllPostsHandler)
		engine.ServeHTTP(w, req)

		expectedResult := test.out

		got := w.Code

		// assert that the expectations were met
		assert.Equal(t, expectedResult, got)
	}
}
