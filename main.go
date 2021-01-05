package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

type user struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

const seq = 1

func main() {
	e := echo.New()
	e.PUT("/create", createUser)
}

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)

}
