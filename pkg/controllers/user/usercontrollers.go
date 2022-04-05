package usercontrollers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	current_page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	users, code, err := controller.GetAllUsersService(current_page)
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, users)
}

func (controllers *UserController) GetUserHandler(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userid"))

	user, code, err := controllers.GetUserService(userid)

	if err != nil {
		fmt.Print(err.Error())
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	userid, _ := strconv.Atoi(c.Param("userid"))

	code, err := controllers.UpdateUserService(userid, data)

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
	userid, _ := strconv.Atoi(c.Param("userid"))

	code, err := controllers.DeleteUserService(userid)

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
