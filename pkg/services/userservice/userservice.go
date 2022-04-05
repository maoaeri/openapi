package userservice

import (
	"errors"
	"net/http"

	"github.com/maoaeri/openapi/pkg/helper"
	"github.com/maoaeri/openapi/pkg/model"
	"github.com/maoaeri/openapi/pkg/repositories/userrepo"
	"gorm.io/gorm"
)

type UserService struct {
	userrepo.IUserRepository
}

/*func NewService(r Repository) Service {
	return Service{repo: r}
}*/

type IUserService interface {
	SignUpService(user *model.User) (code int, err error)
	GetAllUsersService(page int) (users []model.User, code int, err error)
	GetUserService(userid int) (user *model.User, code int, err error)
	UpdateUserService(userid int, data map[string]interface{}) (code int, err error)
	DeleteUserService(userid int) (code int, err error)
	DeleteAllUsersService() (code int, err error)
}

func (service *UserService) SignUpService(user *model.User) (code int, err error) {
	switch "" {
	case user.Username:
		err := errors.New("Username cannot be blank.")
		return http.StatusBadRequest, err
	case user.Email:
		err := errors.New("Email cannot be blank.")
		return http.StatusBadRequest, err
	case user.Password:
		err := errors.New("Password cannot be blank.")
		return http.StatusBadRequest, err
	}

	if service.CheckDuplicateEmail(user.Email) == true {
		err := errors.New("Email already used.")
		return http.StatusBadRequest, err
	}

	user.Password = helper.GenerateHash(user.Password)
	user.Role = "user"

	err = service.CreateUser(user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (service *UserService) GetAllUsersService(page int) (users []model.User, code int, err error) {
	users, err = service.GetAllUsers(page)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("There is no such user.")
		return nil, http.StatusBadRequest, err
	} else if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return users, http.StatusOK, nil
}

func (service *UserService) GetUserService(userid int) (user *model.User, code int, err error) {
	user, err = service.GetUser(userid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("There is no such user.")
		return nil, http.StatusBadRequest, err
	} else if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}

func (service *UserService) UpdateUserService(userid int, data map[string]interface{}) (code int, err error) {
	if userid != data["userid"].(int) {
		err = errors.New("Wrong id")
		return http.StatusBadRequest, err
	}

	err = service.UpdateUser(userid, data)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil

}

func (service *UserService) DeleteUserService(userid int) (code int, err error) {
	err = service.DeleteUser(userid)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil

}

func (service *UserService) DeleteAllUsersService() (code int, err error) {
	err = service.DeleteAllUsers()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
