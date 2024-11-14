package task

import "errors"

type TaskRepository struct {
    tasks []Task
    nextID int
}

func NewTaskRepository() *TaskRepository {
    return &TaskRepository{nextID: 1}
}

func (r *TaskRepository) AddTask(name string) Task {
    task := Task{ID: r.nextID, Name: name, Done: false}
    r.tasks = append(r.tasks, task)
    r.nextID++
    return task
}

func (r *TaskRepository) ListTasks() []Task {
    return r.tasks
}

func (r *TaskRepository) CompleteTask(id int) error {
    for i, task := range r.tasks {
        if task.ID == id {
            r.tasks[i].Done = true
            return nil
        }
    }
    return errors.New("task not found")
}

