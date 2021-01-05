package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type user struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

const seq = 1

var users = map[int]*user{}

func main() {
	e := echo.New()
	e.POST("/users/create", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)

	e.Logger.Fatal(e.Start(":3000"))
}

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	return c.JSON(http.StatusOK, u)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	iid, _ := strconv.Atoi(id)

	return c.JSON(http.StatusOK, users[iid])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id := c.Param("id")
	iid, _ := strconv.Atoi(id)
	users[iid].Name = u.Name
	return c.JSON(http.StatusOK, users[iid])
}
