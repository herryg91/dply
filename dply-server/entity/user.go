package entity

type UserType string

const (
	UserType_Admin UserType = "admin"
	UserType_User  UserType = "user"
)

type User struct {
	Id       int      `json:"id"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	UserType UserType `json:"usertype"`
	Name     string   `json:"name"`
	Token    string   `json:"token"`
}
