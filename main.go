package main

import (
	"covid_api/src"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/stateCases", src.GetStateName)
	e.GET("/saveCases", src.SaveCovidData)
	e.Logger.Fatal(e.Start(":1323"))
}
