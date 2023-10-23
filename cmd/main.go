package main

import (
	"log"

	"github.com/kalougata/go-take-out/cmd/wire"
)

func main() {
	app, cleanup, err := wire.NewApp()

	if err != nil {
		log.Fatal(err)
	}

	app.Listen(":3000")

	defer cleanup()
}
