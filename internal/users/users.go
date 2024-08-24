package users

const (
	USER_ADMIN = 1
	USER_DEFAULT = 2
)

type User struct {
	Type int
	Name string
	Email string
	Key string
	SignUpTime string
}
