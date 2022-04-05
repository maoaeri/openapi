package userrepo

import (
	"errors"
	"fmt"

	"github.com/maoaeri/openapi/pkg/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CheckDuplicateEmail(email string) bool
	CreateUser(user *model.User) error
	GetAllUsers(page int) (users []model.User, err error)
	GetUser(userid int) (user *model.User, err error)
	UpdateUser(userid int, data map[string]interface{}) error
	DeleteUser(userid int) error
	DeleteAllUsers() error
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (userrepo *UserRepo) CheckDuplicateEmail(email string) bool {

	var dbuser model.User
	result := userrepo.DB.Where("email = ?", email).First(&dbuser)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func (userrepo *UserRepo) CreateUser(user *model.User) error {

	result := userrepo.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//page starts with 1
func (userrepo *UserRepo) GetAllUsers(page int) (users []model.User, err error) {

	result := userrepo.DB.Limit(10).Offset((page - 1) * 10).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (userrepo *UserRepo) GetUser(userid int) (user *model.User, err error) {

	result := userrepo.DB.First(&user, "user_id = ?", userid)
	if result.Error != nil {
		fmt.Println("Error in fetching user")
		return user, result.Error
	}
	return user, nil
}

func (userrepo *UserRepo) UpdateUser(userid int, data map[string]interface{}) error {

	user, _ := userrepo.GetUser(userid)
	result := userrepo.DB.Model(&user).Where("user_id = ?", userid).Updates(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userrepo *UserRepo) DeleteUser(userid int) error {

	var user *model.User
	user, _ = userrepo.GetUser(userid)

	result := userrepo.DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userrepo *UserRepo) DeleteAllUsers() error {

	var users []model.User
	userrepo.DB.Find(&users)

	result := userrepo.DB.Delete(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
