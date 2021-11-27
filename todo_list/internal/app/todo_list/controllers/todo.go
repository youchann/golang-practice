package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/infrastructures"
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/repositories"
)

type TodoController struct {
	todoRepository *repositories.TodoRepository
}

func NewTodoController() *TodoController {
	db := infrastructures.GetDbConnection()
	todoRepository := repositories.NewTodoRepository(db)
	return &TodoController{todoRepository: todoRepository}
}

func (tc *TodoController) Index(c echo.Context) error {
	todos, err := tc.todoRepository.GetAll()
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(todos)
	return c.JSON(http.StatusOK, "hoge")
}
