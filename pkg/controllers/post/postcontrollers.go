package postcontrollers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt_handler "github.com/maoaeri/openapi/pkg"
	"github.com/maoaeri/openapi/pkg/model"
	"github.com/maoaeri/openapi/pkg/services/postservice"
)

type PostController struct {
	postservice.IPostService
}

func (controllers *PostController) CreatePostHandler(c *gin.Context) {

	authmiddleware := jwt_handler.JwtHandler()

	var post *model.Post
	err := c.BindJSON(&post)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
	}

	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	post.UserID = int(claims["userid"].(float64))

	code, err := controllers.CreatePostService(post)
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(code, gin.H{
			"message": "Post created successfully.",
			"info":    post,
		})
	}
}

func (controllers *PostController) UpdatePostHandler(c *gin.Context) {

	postid, _ := strconv.Atoi(c.Param("postid"))

	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
	}

	authmiddleware := jwt_handler.JwtHandler()

	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}
	userid_token := int(claims["userid"].(float64))

	code, err := controllers.UpdatePostService(postid, userid_token, data)

	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": "An error ocurred",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
	})
	return
}

func (controllers *PostController) DeletePostHandler(c *gin.Context) {

	postid, _ := strconv.Atoi(c.Param("postid"))

	authmiddleware := jwt_handler.JwtHandler()

	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}
	userid_token := int(claims["userid"].(float64))

	code, err := controllers.DeletePostService(postid, userid_token)

	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
	return
}

func (controllers *PostController) DeleteAllPostsHandler(c *gin.Context) {
	code, err := controllers.DeleteAllPostsService()
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "All users deleted.",
		})
	}
}

func (controllers *PostController) GetPostHandler(c *gin.Context) {

	var post *model.Post

	postid, _ := strconv.Atoi(c.Param("postid"))

	post, code, err := controllers.GetPostService(postid)
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"info": post,
	})
	return
}

func (controllers *PostController) GetAllPostsHandler(c *gin.Context) {
	current_page, _ := strconv.Atoi(c.Query("page"))

	posts, code, err := controllers.GetAllPostsService(current_page)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, posts)
	return
}
