package postservice

import (
	"errors"
	"net/http"
	"testing"

	mocks "github.com/maoaeri/openapi/mocks/pkg/repositories/postrepo"
	"github.com/maoaeri/openapi/pkg/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var errGeneral = errors.New("general error.")

func TestCreatePost(t *testing.T) {
	type Test struct {
		inPost *model.Post
		out    int
		err    error
	}

	postTest := []Test{
		{&model.Post{Content: ""}, http.StatusBadRequest, errGeneral},
		{&model.Post{Content: "ahihi"}, http.StatusInternalServerError, errGeneral},
		{&model.Post{Content: "ahihi"}, http.StatusCreated, nil},
	}
	for _, test := range postTest {
		// create an instance of our test object
		postRepo := new(mocks.IPostRepository)

		//set up expectations
		postRepo.On("CreatePost", test.inPost).Return(test.err)
		postService := PostService{
			postRepo,
		}

		got, _ := postService.CreatePostService(test.inPost)
		assert.Equal(t, test.out, got)
	}
}

func TestUpdatePost(t *testing.T) {
	type Test struct {
		data map[string]interface{}
		out  int
		err  error
	}

	postTest := []Test{
		{map[string]interface{}{"Content": "ahihi"}, http.StatusInternalServerError, errGeneral},
		{map[string]interface{}{"Content": "ahihi"}, http.StatusOK, nil},
	}

	for _, test := range postTest {
		// create an instance of our test object
		postRepo := new(mocks.IPostRepository)

		//set up expectations
		postRepo.On("UpdatePost", 0, test.data).Return(test.err)
		postService := PostService{
			postRepo,
		}

		got, _ := postService.UpdatePostService(0, test.data)
		assert.Equal(t, test.out, got)
	}
}

func TestDeletePost(t *testing.T) {
	type Test struct {
		out int
		err error
	}

	postTest := []Test{
		{http.StatusInternalServerError, errGeneral},
		{http.StatusOK, nil},
	}

	for _, test := range postTest {
		// create an instance of our test object
		postRepo := new(mocks.IPostRepository)

		//set up expectations
		//postid_param = 0 for all test
		postRepo.On("DeletePost", 0).Return(test.err)
		postService := PostService{
			postRepo,
		}

		got, _ := postService.DeletePostService(0)
		assert.Equal(t, test.out, got)
	}
}

func TestGetPost(t *testing.T) {
	type Test struct {
		out int
		err error
	}

	postTest := []Test{
		{http.StatusBadRequest, gorm.ErrRecordNotFound},
		{http.StatusInternalServerError, errGeneral},
		{http.StatusOK, nil},
	}

	var postData = &model.Post{
		PostID:  0,
		Content: "ahihi",
		UserID:  0,
	}
	for _, test := range postTest {
		// create an instance of our test object
		postRepo := new(mocks.IPostRepository)

		//set up expectations
		postRepo.On("GetPost", 0).Return(postData, test.err) //postid_param = 0 for all test
		postService := PostService{
			postRepo,
		}

		_, got, _ := postService.GetPostService(0)
		assert.Equal(t, test.out, got)
	}
}

func TestGetAllPosts(t *testing.T) {
	type Test struct {
		out int
		err error
	}

	postTest := []Test{
		{http.StatusBadRequest, gorm.ErrRecordNotFound},
		{http.StatusInternalServerError, errGeneral},
		{http.StatusOK, nil},
	}

	var postsData = []model.Post{}
	for _, test := range postTest {
		// create an instance of our test object
		postRepo := new(mocks.IPostRepository)

		//set up expectations
		postRepo.On("GetAllPosts", 0).Return(postsData, test.err) //postid_param = 0 for all test
		postService := PostService{
			postRepo,
		}

		_, got, _ := postService.GetAllPostsService(0)
		assert.Equal(t, test.out, got)
	}
}

func TestDeleteAllPosts(t *testing.T) {
	type Test struct {
		out int
		err error
	}

	postTest := []Test{
		{http.StatusInternalServerError, errGeneral},
		{http.StatusOK, nil},
	}

	for _, test := range postTest {
		// create an instance of our test object
		postRepo := new(mocks.IPostRepository)

		//set up expectations
		postRepo.On("DeleteAllPosts").Return(test.err) //postid_param = 0 for all test
		postService := PostService{
			postRepo,
		}

		got, _ := postService.DeleteAllPostsService()
		assert.Equal(t, test.out, got)
	}
}
