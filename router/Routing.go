package router

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"github.com/labstack/echo/middleware"
	config "github.com/joho/godotenv"
	"github.com/lab/ksp-api/component"
	cont "github.com/lab/ksp-api/module/welcome/controller"
)

func Routing()  {

	err := config.Load(".env")
	if err != nil {
		fmt.Println(echo.NewHTTPError(http.StatusInternalServerError, ".env is not loaded properly"))
		os.Exit(2)
	}

	// initiate the host variable
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	app.Binder = &component.CustomBinder{}
	app.HTTPErrorHandler = component.AppHTTPErrorHandler

	//welcome message
	app.GET("/", cont.Welcome)


	// Default Port
	app.Start(os.Getenv("DEFAULT_PORT"))

}
