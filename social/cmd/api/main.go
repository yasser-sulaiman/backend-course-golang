package main

import (
	"log"
	"social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	log.Fatal(app.run(app.mount()))
}
