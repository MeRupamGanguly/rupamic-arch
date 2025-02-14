package domain

type User struct {
	Id       string
	Name     string
	Password string
	Email    string
	Roles    []string
}
