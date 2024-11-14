package task

type TaskService struct {
    repo *TaskRepository
}

func NewTaskService(repo *TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(name string) Task {
    return s.repo.AddTask(name)
}

func (s *TaskService) ListTasks() []Task {
    return s.repo.ListTasks()
}

func (s *TaskService) CompleteTask(id int) error {
    return s.repo.CompleteTask(id)
}

