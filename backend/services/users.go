package services

type CreateNewUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

type UserResposne struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type UserService interface {
	CreateUserAccount(newUser CreateNewUserRequest) (*UserResposne, error)
	GetUser(id int) (*UserResposne, error)
	SignIn(email string, password string) (string, error)
}
