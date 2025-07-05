package main

import (
	"log"

	"github.com/LootNex/TestTask_WorkMate/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден или не удалось загрузить")
	}

	err = server.StartServer()

	if err != nil {
		log.Println(err)
	}

}
