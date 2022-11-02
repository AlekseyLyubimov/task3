package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	addr := "localhost:8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/task3/C", CreateHandler)
	mux.HandleFunc("/task3/R", ReadHandler)
	mux.HandleFunc("/task3/U", UpdateHandler)
	mux.HandleFunc("/task3/D", DeleteHandler)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, LoggerMiddleware(mux)))
}

func LoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("Logger middleware says: %s %s %v", r.Method, r.URL.Path, time.Now().Format(time.StampMilli))
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

}

func ReadHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

}

type Blog struct {
	ID      int
	Login   string
	Blocked bool
}
