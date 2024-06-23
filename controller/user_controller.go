package controller

import (
	"graded-challange-1-ysmnababan/helper"
	"graded-challange-1-ysmnababan/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UR repository.UserRepo
}

func (c *UserController) GetAllUsers(e echo.Context) error {
	res, err := c.UR.GetAllUser()
	if err != nil {
		return helper.ParseError(err, e)
	}
	return e.JSON(http.StatusOK, map[string]interface{}{"Message:": "All users list", "Data": res})
}
