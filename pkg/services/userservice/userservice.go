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
	GetUserService(email_param string, email_token string) (user *model.User, code int, err error)
	UpdateUserService(email_param string, email_token string, data map[string]interface{}) (code int, err error)
	DeleteUserService(email_param string, email_token string) (code int, err error)
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
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return users, http.StatusOK, nil
}

func (service *UserService) GetUserService(email_param string, email_token string) (user *model.User, code int, err error) {
	if email_param == email_token {
		user, err = service.GetUser(email_param)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("There is no such user.")
			return nil, http.StatusBadRequest, err
		} else if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		return user, http.StatusOK, nil
	} else {
		err = errors.New("You cannot get other user's information.")
		return nil, http.StatusBadRequest, err
	}
}

func (service *UserService) UpdateUserService(email_param string, email_token string, data map[string]interface{}) (code int, err error) {
	if email_param == email_token {
		if email_param != data["email"] {
			err = errors.New("Wrong email")
			return http.StatusBadRequest, err
		}

		err = service.UpdateUser(email_param, data)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusOK, nil
	} else {
		err = errors.New("You cannot update other user's information.")
		return http.StatusBadRequest, err
	}
}

func (service *UserService) DeleteUserService(email_param string, email_token string) (code int, err error) {
	if email_param == email_token {
		err = service.DeleteUser(email_param)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusOK, nil
	} else {
		err = errors.New("You cannot delete other user's information.")
		return http.StatusBadRequest, err
	}
}

func (service *UserService) DeleteAllUsersService() (code int, err error) {
	err = service.DeleteAllUsers()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
