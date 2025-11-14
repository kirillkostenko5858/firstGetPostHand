package services

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) error
	GetAllTask() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(task Task) error
	DeleteTaskByID(id string) error
}

type tskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &tskRepository{db: db}
}

func (r *tskRepository) CreateTask(task Task) error {
	return r.db.Create(&task).Error
}
func (r *tskRepository) GetAllTask() ([]Task, error) {
	var rB []Task
	err := r.db.Find(&rB).Error
	return rB, err
}

func (r *tskRepository) GetTaskByID(id string) (Task, error) {
	var rB Task
	err := r.db.First(&rB, "id = ?", id).Error
	return rB, err
}

func (r *tskRepository) UpdateTask(task Task) error {
	return r.db.Save(&task).Error
}

func (r *tskRepository) DeleteTaskByID(id string) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
