package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LootNex/TestTask_WorkMate/internal/model"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type MockTaskManager struct {
	ProcessTaskFunc func(task *model.Task)
	CreateTaskFunc  func() (string, error)
	DeleteTaskFunc  func(id string) error
	GetResultFunc   func(id string) (*model.Task, error)
}

func (m MockTaskManager) ProcessTask(task *model.Task) {
	m.ProcessTaskFunc(task)
}

func (m MockTaskManager) CreateTask() (string, error) {
	return m.CreateTaskFunc()
}

func (m MockTaskManager) DeleteTask(id string) error {
	return m.DeleteTaskFunc(id)
}

func (m MockTaskManager) GetResult(id string) (*model.Task, error) {
	return m.GetResultFunc(id)
}

func TestTaskCreate_Success(t *testing.T) {

	mock := MockTaskManager{
		CreateTaskFunc: func() (string, error) {
			return "", nil
		},
		ProcessTaskFunc: func(task *model.Task) {
		},
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	h := NewHandler(mock, logger)

	req := httptest.NewRequest(http.MethodPost, "/orders", nil)

	w := httptest.NewRecorder()

	h.TaskCreate(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

}

func TestTaskCreate_InternalServerError(t *testing.T) {

	mock := MockTaskManager{
		CreateTaskFunc: func() (string, error) {
			return "", fmt.Errorf("cannot create task")
		},
		ProcessTaskFunc: func(task *model.Task) {
		},
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	h := NewHandler(mock, logger)

	req := httptest.NewRequest(http.MethodPost, "/orders", nil)

	w := httptest.NewRecorder()

	h.TaskCreate(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

}

func TestTaskDelete_Success(t *testing.T) {

	mock := MockTaskManager{
		DeleteTaskFunc: func(id string) error {
			return nil
		},
	}

	r := httptest.NewRequest(http.MethodDelete, "/orders/12345", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "12345"})
	w := httptest.NewRecorder()

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	h := NewHandler(mock, logger)

	h.TaskDelete(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("unexpected status %v, expected %v", res.StatusCode, http.StatusOK)
	}

}

func TestTaskDelete_BadRequest(t *testing.T) {

	mock := MockTaskManager{
		DeleteTaskFunc: func(id string) error {
			return fmt.Errorf("now found")
		},
	}

	r := httptest.NewRequest(http.MethodDelete, "/orders/12345", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "12345"})
	w := httptest.NewRecorder()

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	h := NewHandler(mock, logger)

	h.TaskDelete(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("unexpected status %v, expected %v", res.StatusCode, http.StatusOK)
	}

}

func TestGetResult_Success(t *testing.T) {

	task := model.Task{
		Status:    "in process",
		StartTime: time.Now(),
		Duration:  time.Since(time.Now()),
	}

	mock := MockTaskManager{
		GetResultFunc: func(id string) (*model.Task, error) {
			return &task, nil
		},
	}

	r := httptest.NewRequest(http.MethodGet, "/orders/12345", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "12345"})
	w := httptest.NewRecorder()

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	h := NewHandler(mock, logger)

	h.GetResult(w, r)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("unexpected status %v, expected %v", res.StatusCode, http.StatusOK)
	}

}

func TestGetResult_BadRequest(t *testing.T) {

	mock := MockTaskManager{
		GetResultFunc: func(id string) (*model.Task, error) {
			return nil, fmt.Errorf("cannot get result task")
		},
	}

	r := httptest.NewRequest(http.MethodGet, "/orders/12345", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "12345"})
	w := httptest.NewRecorder()

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	h := NewHandler(mock, logger)

	h.GetResult(w, r)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("unexpected status %v, expected %v", res.StatusCode, http.StatusOK)
	}

}
