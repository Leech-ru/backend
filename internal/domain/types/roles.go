package types

type Role int

const (
	RoleUser Role = iota
	RoleModerator
	RoleAdmin
)

func (r Role) String() string {
	switch r {
	case RoleUser:
		return "User"
	case RoleModerator:
		return "Moderator"
	case RoleAdmin:
		return "Admin"
	default:
		return "unknown"
	}
}
