package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LootNex/TestTask_WorkMate/internal/model"
	"github.com/LootNex/TestTask_WorkMate/internal/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Handler struct {
	TaskHandler *service.TaskService
	Log         *zap.Logger
}

func NewHandler(handler *service.TaskService, log *zap.Logger) *Handler {
	return &Handler{
		TaskHandler: handler,
		Log:         log,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	id, err := h.TaskHandler.CreateOrder()
	if err != nil {
		h.Log.Error(fmt.Sprintf("cannot create task, err: %v", err))
		http.Error(w, "try again", http.StatusInternalServerError)
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

func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	err := h.TaskHandler.DeleteOrder(id)
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
	}

	resp := struct {
		ID   string      `json:"id"`
		Task *model.Task `json:"task"`
	}{
		ID:   id,
		Task: task,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		h.Log.Error(fmt.Sprintf("problems with get result encoding, err: %v", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

}
