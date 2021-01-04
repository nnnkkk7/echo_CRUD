package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type user struct {
	ID   string `json: "id"`
	Name string `json: "name"`
}

func main() {

}

func createUser(c echo.Context)
