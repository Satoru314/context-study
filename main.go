package main

import (
	"context-study/handler"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("study")
	e := echo.New()
	e.Use()
	e.GET("/", handler.HelloHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
