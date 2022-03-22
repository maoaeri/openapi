package postservice

import (
	"github.com/maoaeri/openapi/pkg/model"
	"github.com/maoaeri/openapi/pkg/repositories/postrepo"
)

type PostService struct {
	postrepo.IPostRepository
}

type IPostService interface {
	CreatePostService(post *model.Post) error
}

func (service *PostService) CreatePostService(post *model.Post) error {
	err := service.CreatePost(post)
	return err
}
