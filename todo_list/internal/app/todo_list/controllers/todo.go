package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/infrastructures"
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/repositories"
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/schemas/models"
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/schemas/requestMappers"
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
		panic("failed to get todos")
	}
	return c.JSON(http.StatusOK, todos)
}

func (tc *TodoController) Create(c echo.Context) error {
	newTodo := new(requestMappers.TodoCreate)
	if err := c.Bind(newTodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(newTodo); err != nil {
		return err
	}
	todo, err := tc.todoRepository.Create(models.Todo{Name: newTodo.Name, IsDone: newTodo.IsDone})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todo)
}
