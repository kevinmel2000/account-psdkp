package controller

import (
	"net/http"
	"github.com/labstack/echo"
)

func Welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, "asasass")
}