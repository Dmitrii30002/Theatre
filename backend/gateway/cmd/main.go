package main

import (
	"log"

	"gateway/cmd/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
