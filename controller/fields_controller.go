package controller

import (
	"graded-challange-1-ysmnababan/helper"
	"graded-challange-1-ysmnababan/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FieldController struct {
	FR repository.FieldRepo
}

func (c *FieldController) GetAllFields(e echo.Context) error {
	res, err := c.FR.GetAllFields()
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{"Message:": "All field lists", "Data": res})
}
