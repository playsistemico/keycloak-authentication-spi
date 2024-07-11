package main

import (
	"log"

	"backend/internal/router"
)

func main() {
	r := router.New()

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
