package todolist

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/controllers"
)

func Start() {
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())
	todoController := controllers.NewTodoController()
	api := e.Group("/api")
	{
		api.GET("/", todoController.Index)
		// api.POST("/members", controller.PostMembers())
		// api.PUT("/members", controller.PutMembers())
		// api.DELETE("/members", controller.DeleteMembers())
	}
	e.Logger.Fatal(e.Start(":1323"))
}
