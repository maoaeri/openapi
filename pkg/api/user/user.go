package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/maoaeri/openapi/pkg/api/database"
	"github.com/maoaeri/openapi/pkg/helper"
	"github.com/maoaeri/openapi/pkg/model"
)

func CreateUser(user *model.User) error {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	result := connection.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SignUpHandler(c *gin.Context) {

	connection := api.GetDB()
	defer api.CloseDB(connection)

	var user model.User

	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err.Error())
	}

	if user.UserName == "" || user.Email == "" || user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Username, email and password cannot be blank.",
		})
		return
	}

	var dbuser model.User
	result := connection.Where("email = ?", user.Email).First(&dbuser)
	if result.Error.Error() != "record not found" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Email already in use.",
		})
		return
	} else {
		user.Password = helper.GenerateHash(user.Password)
		user.Role = "user"
		if err := CreateUser(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"message": "Sign Up successfully.",
			})
		}
	}
}

func GetAllUsers() (users []model.User) {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	result := connection.Find(&users)
	if result.Error != nil {
		log.Fatalln("Error in fetching users")
		return nil
	}
	return users
}

func GetUser(email string) (user *model.User, err error) {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	result := connection.Find(&user, "email = ?", email)
	if result.Error != nil {
		fmt.Print("Error in fetching user")
		return user, result.Error
	}
	return user, nil
}

func GetUserHandler(c *gin.Context) {
	email := c.Param("email")

	user, err := GetUser(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusFound, user)
	}
}

func UpdateUser(email string, data map[string]interface{}) error {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	user, _ := GetUser(email)
	result := connection.Model(&user).Updates(data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUserHandler(c *gin.Context) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	email := c.Param("email")

	if email != data["email"] {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Wrong email.",
		})
		return
	}

	if err := UpdateUser(email, data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User updated successfully.",
		})
	}

}

func DeleteUser(email string) error {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	var user *model.User
	user, _ = GetUser(email)

	result := connection.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUserHandler(c *gin.Context) {
	email := c.Param("email")

	err := DeleteUser(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "User deleted successfully.",
		})
	}

}

func DeleteAllUsers() error {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	var users []model.User
	users = GetAllUsers()

	result := connection.Delete(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
