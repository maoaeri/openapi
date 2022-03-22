package initapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwt_handler "github.com/maoaeri/openapi/pkg/api"
	"github.com/maoaeri/openapi/pkg/api/post"
	controllers "github.com/maoaeri/openapi/pkg/controllers/user"
	"github.com/maoaeri/openapi/pkg/database"
	"github.com/maoaeri/openapi/pkg/middleware"
	"github.com/maoaeri/openapi/pkg/repositories/userrepo"
	"github.com/maoaeri/openapi/pkg/services/userservice"
)

var router *gin.Engine

func CreateRouter() {
	router = gin.Default()
}

func InitController() *controllers.UserController {
	controller := controllers.UserController{&userservice.UserService{userrepo.NewUserRepo(database.GetDBInstance().DB)}}
	return &controller
}

func InitRouter() {
	controllers := InitController()
	router.Use(middleware.OpenAPIInputValidator())
	user_routes := router.Group("/users")
	{
		user_routes.POST("/login", jwt_handler.JwtHandler().LoginHandler)
		user_routes.POST("/logout", jwt_handler.JwtHandler().LogoutHandler)
		user_routes.POST("/signup", controllers.SignUpHandler)
		user_routes.DELETE("/:email", controllers.DeleteUserHandler)
		user_routes.GET("/:email", controllers.GetUserHandler)
		user_routes.PUT("/:email", controllers.UpdateUserHandler)
		user_routes.GET("", controllers.GetAllUsersHandler)
		user_routes.DELETE("", controllers.DeleteAllUsersHandler)
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
