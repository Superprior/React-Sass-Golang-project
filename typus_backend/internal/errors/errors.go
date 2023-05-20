package errors

import "errors"

// All the errors used for communications between repos, usecases and handlers.
var (
	ErrNonUniqueEmail     = errors.New("non unique email")
	ErrNonUniqueUsername  = errors.New("non unique username")
	ErrNoUserWithId       = errors.New("no user with such id")
	ErrNoUserWithEmail    = errors.New("no user with such email")
	ErrNoUserWithUsername = errors.New("no user with such username")
	ErrNoSampleWithId     = errors.New("no sample with such id")
	ErrServerError        = errors.New("server error")
)
