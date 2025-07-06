package main

import (
	"log"

	"github.com/LootNex/TestTask_WorkMate/internal/server"
)

func main() {

	err := server.StartServer()

	if err != nil {
		log.Println(err)
	}

}
