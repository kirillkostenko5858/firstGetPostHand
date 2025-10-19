package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&requestBody{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

type requestBody struct {
	Task string `json:"task"`
	ID   string `gorm:"primaryKey" json:"id"`
}

func getHandler(c echo.Context) error {
	var rB []requestBody

	if err := db.Find(&rB).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rB)
}

func postHandler(c echo.Context) error {
	var RequestBody requestBody
	if err := c.Bind(&RequestBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	task := requestBody{
		Task: RequestBody.Task,
		ID:   uuid.NewString(),
	}
	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, task)
}

func patchHandler(c echo.Context) error {
	id := c.Param("id")

	var reqBody requestBody
	if err := c.Bind(&reqBody); err != nil {

		return c.JSON(http.StatusBadRequest, err)
	}

	var req requestBody
	if err := db.First(&req, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	req.Task = reqBody.Task
	if err := db.Save(&req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, req)
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&requestBody{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()

	e := echo.New()

	e.Use(middleware.Logger())

	e.POST("/task", postHandler)
	e.GET("/task", getHandler)
	e.PATCH("/task/:id", patchHandler)
	e.DELETE("/task/:id", deleteHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
