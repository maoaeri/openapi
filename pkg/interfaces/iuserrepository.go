package interfaces

import "github.com/maoaeri/openapi/pkg/model"

type IUserRepository interface {
	CreateUser(user *model.User) error
	GetAllUsers() (users []model.User, err error)
	GetUser(email string) (user *model.User, err error)
	UpdateUser(email string, data map[string]interface{}) error
	DeleteUser(email string) error
	DeleteAllUsers() error
}
