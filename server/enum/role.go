package enum

type Role int

const (
	Admin Role = iota + 1
	User
)

func (r Role) RoleToString() string {
	return [...]string{"Admin", "User"}[r-1]
}

func (r Role) EnumIndex() int {
	return int(r)
}
