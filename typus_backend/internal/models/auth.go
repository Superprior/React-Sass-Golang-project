package models

type role int8

// Enum that represents user roles.
// User of each role has all access rights of users with lower role level.
const (
	GUEST role = iota
	USER
	MODERATOR
	ADMIN
)

// Standard user representation model.
type User struct {
	ID           int
	Username     string
	Email        string
	Role         role
	EncryptedPwd string
}

// Credentials used when registering a new user.
type RegisterCredentials struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Credentials used when logging a user in.
// Only email or username needs to be provided.
type LoginCredentials struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
