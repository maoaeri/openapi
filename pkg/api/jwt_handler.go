package jwt_handler

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	api "github.com/maoaeri/openapi/pkg/api/user"
	"github.com/maoaeri/openapi/pkg/helper"
	"github.com/maoaeri/openapi/pkg/model"
)

var identityKey = "email"

func JwtHandler() *jwt.GinJWTMiddleware {
	authmiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "My Realm",
		Key:         []byte(helper.GetEnvVar("SECRETKEY")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				v, _ = api.GetUser(v.Email)
				return jwt.MapClaims{
					"email":    v.Email,
					"userid":   v.UserID,
					"username": v.UserName,
					"role":     v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals model.UserLoginInfo
			if err := c.ShouldBind(&loginVals); err != nil {

				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			user, err := api.GetUser(email)
			if err != nil {
				return nil, err
			}

			if helper.CheckPasswordHash(password, user.Password) {
				return &model.User{
					Email: email,
					Role:  user.Role,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.User); ok && v.Role == "admin" || v.Role == "user" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,

		SendCookie: true,
	})
	if err != nil {
		log.Fatalln("JWT Error:" + err.Error())
	}
	return authmiddleware
}
