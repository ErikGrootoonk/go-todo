package models

import "time"

type Todo struct {
	ID        int64
	Title     string
	Completed bool
	CreatedAt time.Time
}
