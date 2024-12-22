package users

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
