package models

type Role int8

const (
	guest Role = iota
	user
	moderator
	admin
)

type User struct {
	ID           int
	Username     string
	Email        string
	Role         Role
	EncryptedPwd string
}

type RegisterCredentials struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
