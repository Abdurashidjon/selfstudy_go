package models

import "time"

type Content struct {
	Title string
	Body  string
}

type Article struct {
	ID      int
	Content // Promoted fiels
	Author  Person
	CreatedAt *time.Time
}
