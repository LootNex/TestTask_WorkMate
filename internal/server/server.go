package server

import (
	"log"
	"net/http"

	"github.com/LootNex/TestTask_WorkMate/config"
	"github.com/LootNex/TestTask_WorkMate/internal/db"
	"github.com/LootNex/TestTask_WorkMate/internal/handlers"
	"github.com/LootNex/TestTask_WorkMate/internal/logger"
	"github.com/LootNex/TestTask_WorkMate/internal/service"
	"github.com/gorilla/mux"
)

func StartServer() error {

	config, err := config.ConfigLoad()
	if err != nil {
		log.Fatalf("cannot load config err: %v", err)
	}

	db := db.InitDB()

	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("cannot initialize logger: %v", err)
	}
	defer logger.Sync()

	taskservice := service.NewTaskService(db)

	handler := handlers.NewHandler(taskservice, logger)

	r := mux.NewRouter()
	r.HandleFunc("/orders", handler.TaskCreate).Methods("POST")
	r.HandleFunc("/orders/{id}", handler.TaskDelete).Methods("DELETE")
	r.HandleFunc("/orders/{id}", handler.GetResult).Methods("GET")

	logger.Info("Server is running")

	if err := http.ListenAndServe(":"+config.Server.Port, r); err != nil {
		return err
	}

	return nil

}
