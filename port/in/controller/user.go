package controller

import (
	usecase "ddd_prac/domain/use-case"
	"ddd_prac/port/in/controller/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	service usecase.UserService
	e       *echo.Echo
}

func NewUserController(service usecase.UserService, e *echo.Echo) UserController {
	return UserController{service: service, e: e}
}

func (c UserController) Login() {
	c.e.POST("/", func(ctx echo.Context) error {
		var m model.LoginModel
		err := ctx.Bind(&m)
		if len(m.Pw) == 0 || len(m.Email) == 0 {
			return ctx.String(http.StatusBadRequest, "no value pw or emails")
		}
		if err != nil {
			return ctx.String(http.StatusBadRequest, "bad request")
		}
		result, err := c.service.Login(m.Email, m.Pw)

		if err != nil {
			return ctx.String(http.StatusNotFound, err.Error())
		}

		return ctx.String(http.StatusOK, result.ToString())
	})
}

func (c UserController) Register() {
	c.e.POST("/register", func(ctx echo.Context) error {
		var m model.CreateUserModel
		err := ctx.Bind(&m)
		if len(m.Pw) == 0 || len(m.Email) == 0 {
			return ctx.String(http.StatusBadRequest, "no value pw or emails")
		}
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		err = c.service.Register(m.Email, m.Pw)
		if err != nil {
			return ctx.String(http.StatusConflict, err.Error())
		}
		return ctx.String(http.StatusCreated, "어서오세요")
	})
}
