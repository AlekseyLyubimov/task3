package service

import (
	"log"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"encoding/json"

	models "task3/internal/task3/models"
	migrations "task3/migrations"
)

// people on the internet say dlobal variable is better of two ways to pass connection, but if leels kinda janky
var db *gorm.DB

func DBDriverInitialisation() {
	//connecting to db
	var err error
	dsn := "host=localhost user=postgres password=qwerty123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Omsk"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrations.Migrate(db)
}

func LoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("Logger middleware says: %s %s %v", r.Method, r.URL.Path, time.Now().Format(time.StampMilli))
	}
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.First(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(user)
}
