package controllers

import (
	"net/http"

	"../models"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	result := models.GetUser()
	println("foo")
	return c.JSON(http.StatusOK, result)
}
