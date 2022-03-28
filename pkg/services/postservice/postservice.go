package postservice

import (
	"errors"
	"net/http"

	"github.com/maoaeri/openapi/pkg/model"
	"github.com/maoaeri/openapi/pkg/repositories/postrepo"
	"gorm.io/gorm"
)

type PostService struct {
	postrepo.IPostRepository
}

type IPostService interface {
	CreatePostService(post *model.Post) (code int, err error)
	UpdatePostService(postid_param int, userid_token int, data map[string]interface{}) (code int, err error)
	DeletePostService(postid_param int, userid_token int) (code int, err error)
	GetPostService(postid int) (post *model.Post, code int, err error)
	GetAllPostsService(page int) (posts []model.Post, code int, err error)
	DeleteAllPostsService() (code int, err error)
}

func (service *PostService) CreatePostService(post *model.Post) (code int, err error) {
	if post.Content == "" {
		err = errors.New("Post content cannot be blank.")
		return http.StatusBadRequest, err
	}
	err = service.CreatePost(post)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (service *PostService) UpdatePostService(postid_param int, userid_token int, data map[string]interface{}) (code int, err error) {
	post, err := service.GetPost(postid_param)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if post.UserID == int(userid_token) {

		err = service.UpdatePost(postid_param, data)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusOK, nil
	} else {
		err = errors.New("You cannot update other user's post.")
		return http.StatusBadRequest, err
	}
}

func (service *PostService) DeletePostService(postid_param int, userid_token int) (code int, err error) {
	post, err := service.GetPost(postid_param)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if post.UserID == userid_token {
		err = service.DeletePost(postid_param)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusOK, nil
	} else {
		err = errors.New("You cannot delete other user's post.")
		return http.StatusBadRequest, err
	}
}

//user can read any post
func (service *PostService) GetPostService(postid int) (post *model.Post, code int, err error) {
	post, err = service.GetPost(postid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("No post found.")
			return nil, http.StatusBadRequest, err
		}
		return nil, http.StatusInternalServerError, err
	}

	return post, http.StatusOK, nil
}

func (service *PostService) GetAllPostsService(page int) (posts []model.Post, code int, err error) {
	posts, err = service.GetAllPosts(page)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return posts, http.StatusOK, nil
}

func (service *PostService) DeleteAllPostsService() (code int, err error) {
	err = service.DeleteAllPosts()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
