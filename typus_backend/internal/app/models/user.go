package model

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
