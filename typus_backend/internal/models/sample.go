package models

// Standard code sample representation model.
type Sample struct {
	ID       int
	Title    string
	Content  []string
	Language string
}

// The model of a body that is required to create a new code sample.
type CreateSampleBody struct {
	Title    string   `json:"title"`
	Content  []string `json:"content"`
	Language string   `json:"language"`
}
