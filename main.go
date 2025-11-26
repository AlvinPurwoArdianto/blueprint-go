package main

import (
	"log"

	"blueprint/routes"
)

func main() {
	err := routes.Init()
	if err != nil {
		log.Printf("Error start the server: %s", err)
	}
}
