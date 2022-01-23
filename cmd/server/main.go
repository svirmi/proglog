package main

import (
	"log"
	"proglog/internal/server"
)

const port = ":8080"

func main() {
	srv := server.NewHTTPServer(port)
	log.Fatal(srv.ListenAndServe())
}
