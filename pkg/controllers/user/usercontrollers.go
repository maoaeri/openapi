package usercontrollers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt_handler "github.com/maoaeri/openapi/pkg"
	"github.com/maoaeri/openapi/pkg/model"
	"github.com/maoaeri/openapi/pkg/services/userservice"
)

type UserController struct {
	userservice.IUserService
}

func (controller *UserController) SignUpHandler(c *gin.Context) {

	var user *model.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	code, err := controller.SignUpService(user)
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(code, gin.H{
			"message": "Sign up successfully.",
		})
		return
	}
}

func (controller *UserController) GetAllUsersHandler(c *gin.Context) {
	current_page, _ := strconv.Atoi(c.Query("page"))

	users, code, err := controller.GetAllUsersService(current_page)
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, users)
	return
}

func (controllers *UserController) GetUserHandler(c *gin.Context) {
	email_param := c.Param("email")

	authmiddleware := jwt_handler.JwtHandler()
	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	email_token := claims["email"].(string)
	user, code, err := controllers.GetUserService(email_param, email_token)

	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, user)
	return
}

func (controllers *UserController) UpdateUserHandler(c *gin.Context) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println(err.Error())
	}
	email_param := c.Param("email")

	authmiddleware := jwt_handler.JwtHandler()
	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}
	email_token := claims["email"].(string)

	code, err := controllers.UpdateUserService(email_param, email_token, data)

	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, gin.H{
		"message": "User updated successfully.",
	})
	return
}

func (controllers *UserController) DeleteUserHandler(c *gin.Context) {
	email_param := c.Param("email")

	authmiddleware := jwt_handler.JwtHandler()
	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	email_token := claims["email"].(string)

	code, err := controllers.DeleteUserService(email_param, email_token)

	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, gin.H{
		"message": "User deleted successfully.",
	})
	return
}

func (controllers *UserController) DeleteAllUsersHandler(c *gin.Context) {
	code, err := controllers.DeleteAllUsersService()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(code, gin.H{
			"message": "An error ocurred",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "All users deleted.",
		})
	}
}
