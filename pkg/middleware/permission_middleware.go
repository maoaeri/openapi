package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_handler "github.com/maoaeri/openapi/pkg"
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
func Rejected(c *gin.Context, role string) bool {
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
		if Rejected(c, role) {
			PermissionDenied(c)
			return
		}
		c.Next()
	}
}
