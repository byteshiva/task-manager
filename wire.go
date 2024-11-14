// +build wireinject

package main

import (
    "github.com/byteshiva/task-manager/internal/task"
    "github.com/google/wire"
)

func InitializeTaskService() *task.TaskService {
    wire.Build(task.NewTaskRepository, task.NewTaskService)
    return &task.TaskService{}
}

