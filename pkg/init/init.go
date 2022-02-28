package initapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwt_handler "github.com/maoaeri/openapi/pkg/api"
	"github.com/maoaeri/openapi/pkg/api/post"
	"github.com/maoaeri/openapi/pkg/api/user"
	"github.com/maoaeri/openapi/pkg/middleware"
)

var router *gin.Engine

func CreateRouter() {
	router = gin.Default()
}

func InitRouter() {
	router.Use(middleware.OpenAPIInputValidator())
	user_routes := router.Group("/users")
	{
		user_routes.POST("/login", jwt_handler.JwtHandler().LoginHandler)
		user_routes.POST("/logout", jwt_handler.JwtHandler().LogoutHandler)
		user_routes.POST("/signup", user.SignUpHandler)
		user_routes.DELETE("/:email", user.DeleteUserHandler)
		user_routes.GET("/:email", user.GetUserHandler)
		user_routes.PUT("/:email", user.UpdateUserHandler)
	}
	post_routes := router.Group("/posts")
	{
		post_routes.POST("", post.CreatePostHandler)
		post_routes.PUT("/:postid", post.UpdatePostHandler)
		post_routes.DELETE("/:postid", post.DeletePostHandler)
		post_routes.GET("/:postid", post.GetPostHandler)
		post_routes.GET("", post.GetAllPostsHandler)
	}
}

func Run() {
	CreateRouter()
	InitRouter()
	http.ListenAndServe(":8080", router)
}
