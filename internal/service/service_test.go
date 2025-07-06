package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/LootNex/TestTask_WorkMate/internal/model"
)

type MockDataBase struct {
	CreateFunc func(id string, task *model.Task)
	DeleteFunc func(id string)
	GetFunc    func(id string) (*model.Task, error)
}

func (m MockDataBase) Create(id string, task *model.Task) {
	m.CreateFunc(id, task)
}

func (m MockDataBase) Delete(id string) {
	m.DeleteFunc(id)
}

func (m MockDataBase) Get(id string) (*model.Task, error) {

	return m.GetFunc(id)

}

func TestDeleteTask_Success(t *testing.T) {

	_, cancel := context.WithCancel(context.Background())

	task := model.Task{
		Cancel: cancel,
	}

	mock := MockDataBase{
		GetFunc: func(id string) (*model.Task, error) {
			return &task, nil
		},
		DeleteFunc: func(id string) {},
	}

	ser := NewTaskService(mock)

	err := ser.DeleteTask("")
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}

}

func TestDeleteTask_Fail(t *testing.T) {

	mock := MockDataBase{
		GetFunc: func(id string) (*model.Task, error) {
			return nil, fmt.Errorf("cannot get")
		},
		DeleteFunc: func(id string) {},
	}

	ser := NewTaskService(mock)

	err := ser.DeleteTask("")
	if err.Error() != "cannot get" {
		t.Errorf("unexpected err %v, expected: %v", err, fmt.Errorf("cannot get"))
	}

}

func TestGetResult_Success(t *testing.T) {
	task := model.Task{
		StartTime: time.Now(),
		Duration:  0,
	}
	mock := MockDataBase{GetFunc: func(id string) (*model.Task, error) {
		return &task, nil
	}}

	ser := NewTaskService(mock)

	_, err := ser.GetResult("")
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}
}
func TestGetResult_Fail(t *testing.T) {

	mock := MockDataBase{
		GetFunc: func(id string) (*model.Task, error) {
			return nil, fmt.Errorf("cannot get")
		}}

	ser := NewTaskService(mock)

	_, err := ser.GetResult("")
	if err.Error() != "cannot get" {
		t.Errorf("unexpected err %v, expected: %v", err, fmt.Errorf("cannot get"))
	}
}
