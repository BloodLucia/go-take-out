package main

import (
	"log"

	"github.com/kalougata/go-take-out/cmd/wire"
)

func main() {
	a, f, err := wire.NewApp()

	if err != nil {
		log.Fatal(err)
	}

	a.Listen(":3000")

	defer f()

}
