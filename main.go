package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type requestBody struct {
	Task string `json:"task"`
	ID   string `json:"id"`
}

var rB = []requestBody{}

func postHandler(c echo.Context) error {
	var RequestBody requestBody
	if err := c.Bind(&RequestBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	task := requestBody{
		Task: RequestBody.Task,
		ID:   uuid.NewString(),
	}

	rB = append(rB, task)
	return c.JSON(http.StatusCreated, task)
}

func getHandler(c echo.Context) error {

	if len(rB) == 0 {
		return c.String(http.StatusOK, "Empty request")
	}

	return c.String(http.StatusOK, "Hello "+rB[0].Task)
}

func patchHandler(c echo.Context) error {
	id := c.Param("id")

	var reqBody requestBody
	if err := c.Bind(&reqBody); err != nil {

		return c.JSON(http.StatusBadRequest, err)
	}

	for i, taskItem := range rB {
		if taskItem.ID == id {
			rB[i].Task = reqBody.Task

			return c.JSON(http.StatusOK, rB[i])
		}
	}
	return c.JSON(http.StatusNotFound, "Task not found")
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	for i, taskItem := range rB {
		if taskItem.ID == id {
			rB = append(rB[:i], rB[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusBadRequest, "Task not found")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.POST("/task", postHandler)
	e.GET("/task", getHandler)
	e.PATCH("/task/:id", patchHandler)
	e.DELETE("/task/:id", deleteHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
