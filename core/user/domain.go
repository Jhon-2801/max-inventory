package user

type User struct {
	Id       int
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
