package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	jwt_handler "github.com/maoaeri/openapi/pkg/api"
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
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Email already in use.",
		})
		return
	} else {
		user.Password = helper.GenerateHash(user.Password)
		user.Role = "user"
		if err := CreateUser(&user); err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"message": "Sign Up successfully.",
			})
		}
	}
}

func GetAllUsers() (users []model.User, err error) {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	result := connection.Find(&users)
	if result.Error != nil {
		fmt.Println("Error in fetching users")
		return nil, result.Error
	}
	return users, nil
}

func GetAllUsersHandler(c *gin.Context) {
	authmiddleware := jwt_handler.JwtHandler()
	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	if claims["role"].(string) == "admin" {
		users, err := GetAllUsers()
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
			return
		} else {
			var output []model.User

			current_page, _ := strconv.Atoi(c.Query("page"))
			if len(users) >= (current_page-1)*10+1 {
				if len(users) >= current_page*10 {
					for i := (current_page-1)*10 + 1; i <= current_page*10; i++ {
						output = append(output, users[i])
					}
				} else {
					for i := (current_page-1)*10 + 1; i <= len(users); i++ {
						output = append(output, users[i-1])
					}
				}
				c.JSON(http.StatusAccepted, output)
				return
			}
		}
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot get all users' information.",
		})
	}
}

func GetUser(email string) (user *model.User, err error) {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	result := connection.First(&user, "email = ?", email)
	if result.Error != nil {
		fmt.Println("Error in fetching user")
		return user, result.Error
	}
	return user, nil
}

func GetUserHandler(c *gin.Context) {
	email := c.Param("email")

	authmiddleware := jwt_handler.JwtHandler()
	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	if email == claims["email"].(string) || claims["role"].(string) == "admin" {
		user, err := GetUser(email)
		if err != nil {
			fmt.Println(err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "There is no such user.",
				})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "An error ocurred",
				})
				return
			}

		} else {
			c.JSON(http.StatusFound, user)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot get other user's information.",
		})
	}
}

func UpdateUser(email string, data map[string]interface{}) error {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	user, _ := GetUser(email)
	result := connection.Model(&user).Where("email = ?", email).Updates(data)

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

	authmiddleware := jwt_handler.JwtHandler()
	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	if email == claims["email"].(string) || claims["role"].(string) == "admin" {
		if email != data["email"] {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Wrong email.",
			})
			return
		}

		if err := UpdateUser(email, data); err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"message": "User updated successfully.",
			})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot update other user's information.",
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

	authmiddleware := jwt_handler.JwtHandler()
	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	if email == claims["email"].(string) || claims["role"].(string) == "admin" {
		err := DeleteUser(email)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
		} else {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "User deleted successfully.",
			})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot delete other user's information.",
		})
	}
}

func DeleteAllUsers() error {
	connection := api.GetDB()
	defer api.CloseDB(connection)

	var users []model.User
	users, _ = GetAllUsers()

	result := connection.Delete(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteAllUsersHandler(c *gin.Context) {
	authmiddleware := jwt_handler.JwtHandler()
	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	if claims["role"].(string) == "admin" {
		err := DeleteAllUsers()
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot get all users' information.",
		})
	}
}
