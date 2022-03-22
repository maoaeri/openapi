package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt_handler "github.com/maoaeri/openapi/pkg/api"
)

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

		if claims["role"].(string) == "admin" {
			return
		} else if claims["role"].(string) == "user" {

		}
		c.Next()
	}
}
