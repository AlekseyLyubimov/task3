package main

import (
	"log"
	"net/http"

	service "task3/internal/task3/service"
)

func main() {
	addr := "localhost:8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/task3/R", service.ReadHandler)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, service.LoggerMiddleware(mux)))
}
