package middleware

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-gonic/gin"
)

func OpenAPIInputValidator() gin.HandlerFunc {
	doc, err := openapi3.NewLoader().LoadFromFile("D:/GolangPj/github.com/maoaeri/openapi/docs/openapi3.yaml")
	if err != nil {
		panic(err)
	}

	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		route, pathParams, err := router.FindRoute(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = openapi3filter.ValidateRequest(c.Request.Context(), &openapi3filter.RequestValidationInput{
			Request:    c.Request,
			PathParams: pathParams,
			Route:      route,
			Options: &openapi3filter.Options{
				MultiError: true,
			},
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.Next()
	}
}
