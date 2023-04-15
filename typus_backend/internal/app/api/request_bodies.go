package apiserver

// Body for creating Sample
// @Description Object that stores info about Sample that need to be created
type PostSampleBody struct {
	Title    string `json:"Title"`
	LangSlug string `json:"LangSlug"`
	Content  string `json:"Content"`
}
