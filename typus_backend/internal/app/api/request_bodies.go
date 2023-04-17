package apiserver

// Body for creating Sample
// @Description Object that stores info about Sample that need to be created.
type PostSampleBody struct {
	Title    string `json:"Title"`
	LangSlug string `json:"LangSlug"`
	Content  string `json:"Content"`
}

// Body for registering Users
// @Description Object that contains obligatory info that is needed to create a new User.
type RegisterBody struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// Body for logging a User in
// @Description Object that contains info for logging a User in. Unlike RegisterBody type, only
// @Description Username or Email needs to be provided.
type LoginBody struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
