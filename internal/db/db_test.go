package db

import (
	"context"
	"testing"
	"time"

	"github.com/LootNex/TestTask_WorkMate/internal/model"
)

func TestCreate(t *testing.T) {

	db := InitDB()

	ctx, cancel := context.WithCancel(context.Background())

	task := &model.Task{
		Status:    "created",
		StartTime: time.Now(),
		Duration:  time.Since(time.Now()),
		Ctx:       ctx,
		Cancel:    cancel,
	}
	id := "12345"
	db.Create(id, task)

	if db.storage[id] != task {
		t.Errorf("unexpected result, expected %v, got %v", task, db.storage[id])
	}

}

func TestDelete(t *testing.T) {

	db := InitDB()

	ctx, cancel := context.WithCancel(context.Background())

	task := &model.Task{
		Status:    "created",
		StartTime: time.Now(),
		Duration:  time.Since(time.Now()),
		Ctx:       ctx,
		Cancel:    cancel,
	}
	id := "12345"
	db.Create(id, task)

	if db.storage[id] != task {
		t.Errorf("unexpected result, expected %v, got %v", task, db.storage[id])
	}

	db.Delete(id)

	if len(db.storage) != 0 {
		t.Errorf("db must be empty")
	}
}

func TestGet(t *testing.T) {

	db := InitDB()

	ctx, cancel := context.WithCancel(context.Background())

	task := &model.Task{
		Status:    "created",
		StartTime: time.Now(),
		Duration:  time.Since(time.Now()),
		Ctx:       ctx,
		Cancel:    cancel,
	}
	id := "12345"
	db.Create(id, task)

	if db.storage[id] != task {
		t.Errorf("unexpected result, expected %v, got %v", task, db.storage[id])
	}

	res, err := db.Get(id)
	if err != nil {
		t.Errorf("unexpected error err: %v", err)
	}

	if res != task {
		t.Errorf("unexpected result, expected %v, got %v", task, res)
	}
}
