package main

import (
	"log"
	"webEngineering/pkg/server"
)

func main() {
	log.Println("Start")
	err := server.Start()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
