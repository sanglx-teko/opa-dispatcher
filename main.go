package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sanglx-teko/opa-dispatcher/controller"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/decision/handler", controller.HandleDecisionAPIController)
	e.Logger.Fatal(e.Start(":1323"))
}
