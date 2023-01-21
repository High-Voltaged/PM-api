package types

type UserRole string

const (
	Admin UserRole = "ADMIN"
	User  UserRole = "USER"
)

func (UserRole) Values() (roles []string) {
	for _, s := range []UserRole{Admin, User} {
		roles = append(roles, string(s))
	}
	return
}
