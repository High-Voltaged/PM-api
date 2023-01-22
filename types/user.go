package types

// Struct holding valid user role values
var UserRole = struct {
	ADMIN string
	USER  string
}{
	ADMIN: "admin",
	USER:  "user",
}

var UserRoleValues = []string{UserRole.ADMIN, UserRole.USER}
