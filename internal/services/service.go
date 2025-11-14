package services

type TaskService interface {
	CreateTask(task Task) (Task, error)
	GetAllTask() ([]Task, error)
	GetTaskById(id string) (Task, error)
	UpdateTask(id string, task Task) (Task, error)
	DeleteTask(id string) (Task, error)
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

type taskService struct {
	repo TaskRepository
}

func (s *taskService) CreateTask(task Task) (Task, error) {
	err := s.repo.CreateTask(task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()
}

func (s *taskService) GetTaskById(id string) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id string, task Task) (Task, error) {
	task.ID = id
	err := s.repo.UpdateTask(task)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(id string) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}
	err = s.repo.DeleteTaskByID(id)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
