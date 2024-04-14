package main

import (
	"innoventes-test/internal/server"
	"log"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		log.Default().Printf("server started on port %s", err.Error())
		panic("cannot start server")
	}
}
