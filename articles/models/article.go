package models

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var ArticleList map[string]Article = make(map[string]Article)

// ArticleList = make(map[string]Article)
