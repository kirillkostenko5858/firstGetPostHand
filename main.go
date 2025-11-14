package main

import (
	"First/internal/db"
	"First/internal/handlers"
	"First/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	e := echo.New()

	taskRepo := services.NewTaskRepository(database)
	taskService := services.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.Use(middleware.Logger())

	e.POST("/task", taskHandlers.PostHandler)
	e.GET("/task", taskHandlers.GetHandler)
	e.PATCH("/task/:id", taskHandlers.PatchHandler)
	e.DELETE("/task/:id", taskHandlers.DeleteHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
