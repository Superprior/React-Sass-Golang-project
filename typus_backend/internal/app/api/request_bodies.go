package apiserver

// Body for creating Sample
// @Description Object that stores info about Sample that need to be created
type PostSampleBody struct {
	Title    string `json:"Title"`
	LangSlug string `json:"LangSlug"`
	Content  string `json:"Content"`
}

type RegisterBody struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
