package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/LootNex/TestTask_WorkMate/internal/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Handler struct {
	TaskHandler service.TaskManager
	Log         *zap.Logger
}

func NewHandler(handler service.TaskManager, log *zap.Logger) *Handler {
	return &Handler{
		TaskHandler: handler,
		Log:         log,
	}
}

func (h *Handler) TaskCreate(w http.ResponseWriter, r *http.Request) {

	id, err := h.TaskHandler.CreateTask()
	if err != nil {
		h.Log.Error(fmt.Sprintf("cannot create task, err: %v", err))
		http.Error(w, "try again", http.StatusInternalServerError)
		return
	}

	h.Log.Info(fmt.Sprintf("Task with id: %s created", id))

	resp := struct {
		ID string `json:"id"`
	}{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		h.Log.Error(fmt.Sprintf("problems with create encoding, err: %v", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

}

func (h *Handler) TaskDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	err := h.TaskHandler.DeleteTask(id)
	if err != nil {
		h.Log.Error(fmt.Sprintf("cannot delete task, err: %v", err))
		http.Error(w, "Please, change field id", http.StatusBadRequest)

	} else {
		h.Log.Info(fmt.Sprintf("Task with id: %s deleted", id))
		w.Write([]byte("successful deleted"))
	}

}

func (h *Handler) GetResult(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	task, err := h.TaskHandler.GetResult(id)
	if err != nil {
		h.Log.Error(fmt.Sprintf("cannot get result task, err: %v", err))
		http.Error(w, "Please, change field id", http.StatusBadRequest)
		return
	}

	resp := struct {
		ID        string        `json:"id"`
		Status    string        `json:"status"`
		StartTime time.Time     `json:"starttime"`
		Duration  time.Duration `json:"duration"`
	}{
		ID:        id,
		Status:    task.Status,
		StartTime: task.StartTime,
		Duration:  task.Duration,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		h.Log.Error(fmt.Sprintf("problems with get result encoding, err: %v", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

}
