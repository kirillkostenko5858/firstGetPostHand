package handlers

import (
	"First/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(s services.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetHandler(c echo.Context) error {
	tasks, err := h.service.GetAllTask()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostHandler(c echo.Context) error {
	var request services.Task
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	request.ID = uuid.NewString()
	task, err := h.service.CreateTask(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) PatchHandler(c echo.Context) error {
	id := c.Param("id")

	var request services.Task
	if err := c.Bind(&request); err != nil {

		return c.JSON(http.StatusBadRequest, err)
	}

	updatedTask, err := h.service.UpdateTask(id, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")
	_, err := h.service.GetTaskById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	_, err = h.service.DeleteTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
