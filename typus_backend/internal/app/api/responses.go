package apiserver

// Response with ID object
// @Description Object that is returned in endpoints where ID needs to be returned
type IdResponse struct {
	id int
}

// Response with message
// @Description Object that is returned when the actual body is empty
type MessageResponse struct {
	message string
}
