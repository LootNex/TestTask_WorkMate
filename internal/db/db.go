package db

import (
	"errors"

	"github.com/LootNex/TestTask_WorkMate/internal/model"
)

type DataBase struct {
	storage map[string]*model.Task
}

type StorageTask interface {
	Create(id string, task *model.Task)
	Delete(id string)
	Get(id string) (*model.Task, error)
}

func InitDB() *DataBase {

	return &DataBase{
		storage: make(map[string]*model.Task),
	}
}

func (db *DataBase) Create(id string, task *model.Task) {

	db.storage[id] = task

}

func (db *DataBase) Delete(id string) {

	delete(db.storage, id)

}

func (db *DataBase) Get(id string) (*model.Task, error) {

	task, ok := db.storage[id]

	if !ok {
		return nil, errors.New("no such id in storage")
	}

	return task, nil

}
