package models

import "time"

type Content struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Article struct {
	ID        int `json:"id"`
	Content       // Promoted fiels
	Author    Person `json:"author"`
	CreatedAt *time.Time `json:"create_at"`
}

type ArticleReq struct {
	ID        int `json:"id"`
	Content       // Promoted fiels
	Author    Person `json:"author"`
}

type ArticleCreate struct {
	Content       // Promoted fiels
	Author    Person `json:"author"`
}
