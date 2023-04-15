package apiserver

// Body for creating Sample
// @Description Object that stores info about Sample that need to be created
type PostSampleBody struct {
	Title    string `json:"Title"`
	LangSlug string `json:"LangSlug"`
	Content  string `json:"Content"`
}

// Response with message
// @Description Object that is returned when the actual body is empty
type MessageResponse struct {
	message string
}

// Password object
// @Description Object that needs to be provided for authorizing a user
type PasswordBody struct {
	Pwd string
}

// Response with ID object
// @Description Object that is returned in endpoints where ID needs to be returned
type IdResponse struct {
	id int
}
