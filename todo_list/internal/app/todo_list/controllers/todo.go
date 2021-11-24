package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoController struct{}

func NewTodoController() *TodoController {
	return new(TodoController)
}

func (tc *TodoController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "hoge")
}
