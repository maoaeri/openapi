package post

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt_handler "github.com/maoaeri/openapi/pkg/api"
	"github.com/maoaeri/openapi/pkg/api/database"
	"github.com/maoaeri/openapi/pkg/model"
)

func CreatePostHandler(c *gin.Context) {
	connection := database.GetDB()
	defer database.CloseDB(connection)

	authmiddleware := jwt_handler.JwtHandler()

	var post model.Post
	err := c.BindJSON(&post)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
	}

	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	post.UserID = uint(claims["userid"].(float64))
	post.UserName = claims["username"].(string)

	result := connection.Create(&post)
	if result.Error != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Post created successfully.",
			"info":    post,
		})
	}
}

func UpdatePostHandler(c *gin.Context) {
	connection := database.GetDB()
	defer database.CloseDB(connection)

	var post model.Post

	postid, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}
	connection.Where("post_id = ?", postid).First(&post)

	authmiddleware := jwt_handler.JwtHandler()

	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	if post.UserID == uint(claims["userid"].(float64)) || claims["role"].(string) == "admin" {
		var data map[string]interface{}
		err = c.BindJSON(&data)
		if err != nil {
			fmt.Println(err.Error())
		}
		result := connection.Model(&post).Where("post_id = ?", postid).Updates(data)

		if result.Error != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Post updated successfully",
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot update other user's post.",
		})
	}
}

func DeletePostHandler(c *gin.Context) {
	connection := database.GetDB()
	defer database.CloseDB(connection)

	var post model.Post

	postid, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}
	connection.Where("post_id = ?", postid).First(&post)

	authmiddleware := jwt_handler.JwtHandler()

	claims, err := authmiddleware.GetClaimsFromJWT(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}

	if post.UserID == uint(claims["userid"].(float64)) || claims["role"].(string) == "admin" {
		result := connection.Delete(&post)

		if result.Error != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Post deleted successfully",
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot delete other user's post.",
		})
	}
}

func GetPostHandler(c *gin.Context) {
	connection := database.GetDB()
	defer database.CloseDB(connection)

	var post model.Post

	postid, err := strconv.Atoi(c.Param("postid"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "An error ocurred",
		})
		return
	}
	result := connection.Where("post_id = ?", postid).First(&post)
	if result.Error.Error() == "record not found" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Cannot find post.",
		})
		return
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

	if post.UserID == uint(claims["userid"].(float64)) || claims["role"].(string) == "admin" {
		if result.Error != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"info": post,
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot read other user's post.",
		})
	}
}

func GetAllPostsHandler(c *gin.Context) {
	connection := database.GetDB()
	defer database.CloseDB(connection)

	var posts []model.Post

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
		result := connection.Find(&posts)
		if result.Error != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
			return
		}

		var output []model.Post

		current_page, _ := strconv.Atoi(c.Query("page"))
		if len(posts) >= (current_page-1)*10+1 {
			if len(posts) >= current_page*10 {
				for i := (current_page-1)*10 + 1; i <= current_page*10; i++ {
					output = append(output, posts[i])
				}
			} else {
				for i := (current_page-1)*10 + 1; i <= len(posts); i++ {
					output = append(output, posts[i])
				}
			}
			c.JSON(http.StatusAccepted, output)
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot read all posts.",
		})
	}
}
