package postrepo

import (
	"fmt"

	"github.com/maoaeri/openapi/pkg/model"
	"gorm.io/gorm"
)

type PostRepo struct {
	DB *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		DB: db,
	}
}

type IPostRepository interface {
	CreatePost(post *model.Post) error
	GetAllPosts(page int) (posts []model.Post, err error)
	GetPost(postid int) (post *model.Post, err error)
	UpdatePost(postid int, data map[string]interface{}) error
	DeletePost(postid int) error
	DeleteAllPosts() error
}

func (postrepo *PostRepo) CreatePost(post *model.Post) error {

	result := postrepo.DB.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//page starts with 1
func (postrepo *PostRepo) GetAllPosts(page int) (posts []model.Post, err error) {

	result := postrepo.DB.Limit(10).Offset((page - 1) * 10).Find(&posts)
	if result.Error != nil {
		fmt.Println("Error in fetching Posts")
		return nil, result.Error
	}
	return posts, nil
}

//Get post by id
func (postrepo *PostRepo) GetPost(postid int) (post *model.Post, err error) {

	result := postrepo.DB.Where("post_id = ?", postid).First(&post)
	if result.Error != nil {
		fmt.Println("Error in fetching Post")
		return post, result.Error
	}
	return post, nil
}

func (postrepo *PostRepo) UpdatePost(postid int, data map[string]interface{}) error {

	post, _ := postrepo.GetPost(postid)
	result := postrepo.DB.Model(&post).Where("post_id = ?", postid).Updates(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postrepo *PostRepo) DeletePost(postid int) error {

	var post *model.Post
	post, _ = postrepo.GetPost(postid)

	result := postrepo.DB.Delete(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postrepo *PostRepo) DeleteAllPosts() error {

	var posts []model.Post
	result := postrepo.DB.Find(&posts)

	result = postrepo.DB.Delete(&posts)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
