package types

type Role int

const (
	User Role = iota
	Moderator
	Admin
)
