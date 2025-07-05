package service

import (
	"context"
	"time"

	"github.com/LootNex/TestTask_WorkMate/internal/db"
	"github.com/LootNex/TestTask_WorkMate/internal/model"
	"github.com/google/uuid"
)

type TaskService struct {
	Service db.StorageTask
}

func NewTaskService(service db.StorageTask) *TaskService {
	return &TaskService{
		Service: service,
	}
}

func (t *TaskService) ProcessTask(task *model.Task) {

	task.Status = "in processing"

	select {
	case <-time.After(3 * time.Minute):
	case <-task.Ctx.Done():
		return
	}

	task.Status = "done"
}

func (t *TaskService) CreateOrder() (string, error) {

	id := uuid.New().String()
	status := "created"
	starttime := time.Now()
	duration := time.Since(starttime)
	ctx, cancel := context.WithCancel(context.Background())

	task := &model.Task{
		Status:    status,
		StartTime: starttime,
		Duration:  duration,
		Ctx:       ctx,
		Cancel:    cancel,
	}

	go t.ProcessTask(task)

	t.Service.Create(id, task)

	return id, nil

}

func (t *TaskService) DeleteOrder(id string) error {

	task, err := t.Service.Get(id)

	if err != nil {
		return err
	}

	task.Cancel()

	t.Service.Delete(id)

	return nil

}

func (t *TaskService) GetResult(id string) (*model.Task, error) {

	task, err := t.Service.Get(id)

	if err != nil {
		return nil, err
	}

	task.Duration = time.Duration(time.Since(task.StartTime).Minutes())

	return task, nil

}
