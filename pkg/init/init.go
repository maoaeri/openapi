package initapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwt_handler "github.com/maoaeri/openapi/pkg"
	controllers "github.com/maoaeri/openapi/pkg/controllers"
	postcontrollers "github.com/maoaeri/openapi/pkg/controllers/post"
	usercontrollers "github.com/maoaeri/openapi/pkg/controllers/user"
	"github.com/maoaeri/openapi/pkg/database"
	"github.com/maoaeri/openapi/pkg/middleware"
	"github.com/maoaeri/openapi/pkg/repositories/postrepo"
	"github.com/maoaeri/openapi/pkg/repositories/userrepo"
	"github.com/maoaeri/openapi/pkg/services/postservice"
	"github.com/maoaeri/openapi/pkg/services/userservice"
)

var router *gin.Engine

func CreateRouter() {
	router = gin.Default()
}

func InjectController() *controllers.Controllers {
	db := database.GetDBInstance().DB
	controller := controllers.Controllers{
		UserController: &usercontrollers.UserController{&userservice.UserService{userrepo.NewUserRepo(db)}},
		PostController: &postcontrollers.PostController{&postservice.PostService{postrepo.NewPostRepo(db)}},
	}
	return &controller
}

func InitRouter() {
	controllers := InjectController()
	router.Use(middleware.OpenAPIInputValidator())

	router.POST("/users/login", jwt_handler.JwtHandler().LoginHandler)
	router.POST("/users/logout", jwt_handler.JwtHandler().LogoutHandler)
	router.POST("/users/signup", controllers.UserController.SignUpHandler)

	router.Use(middleware.PermissionMiddleware())
	user_routes := router.Group("/users")
	{
		user_routes.DELETE("/:email", controllers.UserController.DeleteUserHandler)
		user_routes.GET("/:email", controllers.UserController.GetUserHandler)
		user_routes.PUT("/:email", controllers.UserController.UpdateUserHandler)
		user_routes.GET("", controllers.UserController.GetAllUsersHandler)
		user_routes.DELETE("", controllers.UserController.DeleteAllUsersHandler)
	}
	post_routes := router.Group("/posts")
	{
		post_routes.POST("", controllers.PostController.CreatePostHandler)
		post_routes.PUT("/:postid", controllers.PostController.UpdatePostHandler)
		post_routes.DELETE("/:postid", controllers.PostController.DeletePostHandler)
		post_routes.GET("/:postid", controllers.PostController.GetPostHandler)
		post_routes.GET("", controllers.PostController.GetAllPostsHandler)
		post_routes.DELETE("", controllers.PostController.DeleteAllPostsHandler)
	}
}

func Run() {
	CreateRouter()
	InitRouter()
	http.ListenAndServe(":8080", router)
}
