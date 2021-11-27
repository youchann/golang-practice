package todolist

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/controllers"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Start() {
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Validator = &CustomValidator{validator: validator.New()}
	todoController := controllers.NewTodoController()
	api := e.Group("/api")
	{
		api.GET("/", todoController.Index)
		api.POST("/", todoController.Create)
		// api.PUT("/members", controller.PutMembers())
		// api.DELETE("/members", controller.DeleteMembers())
	}
	e.Logger.Fatal(e.Start(":1323"))
}
