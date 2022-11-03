package main

import (
	"log"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"encoding/json"
)

// people on the internet say dlobal variable is better of two ways to pass connection, but if leels kinda janky
var db *gorm.DB

func main() {

	//connecting to db
	var err error
	dsn := "host=localhost user=postgres password=qwerty123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Omsk"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Will automatically create missing columns
	db.AutoMigrate(&User{})
	// Can also be done explicidly
	db.Migrator().AddColumn(&User{}, "LastLogin")

	//testing if it actually works by printing first login
	var user User
	db.First(&user)
	log.Printf(user.Login)

	addr := "localhost:8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/task3/R", ReadHandler)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, LoggerMiddleware(mux)))
}

func LoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("Logger middleware says: %s %s %v", r.Method, r.URL.Path, time.Now().Format(time.StampMilli))
	}
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.First(&user)

	json.NewEncoder(w).Encode(user)
}

type User struct {
	ID        int
	Login     string
	Blocked   bool
	LastLogin time.Time
}
