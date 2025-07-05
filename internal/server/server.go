package server

import (
	"log"
	"net/http"
	"os"

	"github.com/LootNex/TestTask_WorkMate/internal/db"
	"github.com/LootNex/TestTask_WorkMate/internal/handlers"
	"github.com/LootNex/TestTask_WorkMate/internal/logger"
	"github.com/LootNex/TestTask_WorkMate/internal/service"
	"github.com/gorilla/mux"
)

func StartServer() error {

	db := db.InitDB()
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("cannot initialize logger: %v", err)
	}

	taskservice := service.NewTaskService(db)

	handler := handlers.NewHandler(taskservice, logger)

	r := mux.NewRouter()
	r.HandleFunc("/orders", handler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", handler.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/orders/{id}", handler.GetResult).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		logger.Info("PORT не задан, используется порт по умолчанию: 8080")
		port = "8080"
	}

	logger.Info("Server is running")

	if err := http.ListenAndServe(":"+port, r); err != nil {
		return err
	}

	return nil

}
