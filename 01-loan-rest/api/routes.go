package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/ucok-man/h8-p3-preview/01-loan-rest/internal/serializer"
	"github.com/ucok-man/h8-p3-preview/01-loan-rest/internal/validator"
)

func (app *Application) routes() http.Handler {
	router := echo.New()
	router.HTTPErrorHandler = app.httpErrorHandler
	router.JSONSerializer = serializer.JSONSerializer{}
	router.Validator = validator.New()

	root := router.Group("/v1")
	root.Use(app.withRecover())
	root.Use(app.withLogger())

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	users := root.Group("/users")
	{
		users.POST("/register", app.userRegisterHandler)
		users.POST("/login", app.userLoginHandler)
	}

	return router
}
