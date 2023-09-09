package main

import (
	"fmt"
	"webEngineering/pkg/server"
)

func main() {
	err := server.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
