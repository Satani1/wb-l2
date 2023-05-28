package models

import "time"

type Event struct {
	ID    int64     `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"` //RFC3339 format
}
