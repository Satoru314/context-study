package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HelloHandler(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	fmt.Println("user")
	fmt.Println(user)
	return c.String(http.StatusOK, "Hello, World!")
}
