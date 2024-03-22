package userDomain

type User struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func New() User {
	return User{}
}
