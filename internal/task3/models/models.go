package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int
	Login     string
	Blocked   bool
	LastLogin time.Time
}

// is not actually called if good record not found
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.Login == "" {
		log.Printf("Found no record for user with ID = %d", u.ID)
	} else {
		log.Printf("Found good record for user with ID = %d", u.ID)
	}
	return
}
