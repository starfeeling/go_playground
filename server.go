package main

import (
	"net/http"

	"./controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Hello golang world! With mvc echo")
	})
	app.GET("/employees", controllers.GetEmployees)
	app.GET("/users", controllers.GetUsers)

	app.Logger.Fatal(app.Start(":8081"))

}
