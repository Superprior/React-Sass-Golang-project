package models

type Sample struct {
	ID       int
	Title    string
	Content  []string
	Language string
}

type CreateSampleBody struct {
	Title    string   `json:"title"`
	Content  []string `json:"content"`
	Language string   `json:"language"`
}
