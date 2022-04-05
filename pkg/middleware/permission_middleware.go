package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_handler "github.com/maoaeri/openapi/pkg"
	api "github.com/maoaeri/openapi/pkg/database"
	"github.com/maoaeri/openapi/pkg/model"
)

type permInfo struct {
	prefix  string
	methods []string
}

//only user can access
var userPerm = []permInfo{
	{
		prefix:  "/posts",
		methods: []string{"POST", "GET"},
	},
	{
		prefix:  "/posts/",
		methods: []string{"PUT", "GET", "DELETE"},
	},
	{
		prefix:  "/users/",
		methods: []string{"PUT", "GET", "DELETE"},
	},
}

// Rejected checks if a given request should be rejected.
func RejectedRole(c *gin.Context, role string) bool {
	path := c.Request.URL.Path // the path of the url that the user wish to visit
	method := c.Request.Method //the method

	if role == "admin" {
		return false
	}

	if role == "user" {
		for _, info_perm := range userPerm {
			if strings.HasPrefix(path, info_perm.prefix) {
				for _, med := range info_perm.methods {
					if method == med {
						return false
					}
				}
			}
		}
	}

	// Reject
	return true
}

func RejectedInfor(c *gin.Context, role string, userid int) bool {
	path := c.Request.URL.Path

	if role == "admin" {
		return false
	}

	if strings.HasPrefix(path, "/posts/") {
		var post *model.Post
		postid, _ := strconv.Atoi(c.Param("postid"))
		connection := api.GetDB()
		defer api.CloseDB(connection)
		_ = connection.Where("post_id = ?", postid).First(&post)
		if post.UserID == userid {
			return false
		}
	}
	if strings.HasPrefix(path, "/users/") {
		userid_param, _ := strconv.Atoi(c.Param("userid"))
		if userid_param == userid {
			return false
		}
	}
	return true
}

//serve this func if permission is denied
func PermissionDenied(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": "Permission denied.",
	})
	return
}

//middleware
func PermissionMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authmiddleware := jwt_handler.JwtHandler()
		claims, err := authmiddleware.GetClaimsFromJWT(c)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "An error ocurred",
			})
			return
		}
		role := claims["role"].(string)
		userid := int(claims["userid"].(float64))

		if RejectedRole(c, role) {
			PermissionDenied(c)
			return
		}

		if RejectedInfor(c, role, userid) {
			PermissionDenied(c)
			return
		}
		c.Next()
	}
}
