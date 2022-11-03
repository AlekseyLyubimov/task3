package models

import "time"

type User struct {
	ID        int
	Login     string
	Blocked   bool
	LastLogin time.Time
}
