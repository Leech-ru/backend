package types

type Role int

const (
	RoleUser Role = iota
	RoleModerator
	RoleAdmin
	RoleSuperAdmin
)

func (r Role) String() string {
	switch r {
	case RoleUser:
		return "User"
	case RoleModerator:
		return "Moderator"
	case RoleAdmin:
		return "Admin"
	case RoleSuperAdmin:
		return "SuperAdmin"
	default:
		return "unknown"
	}
}
