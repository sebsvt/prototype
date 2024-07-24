package services

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	FirstName   string `json:"firstname"`
	Surname     string `json:"surname"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type TokenData struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"type"`
}

// {
//     "email": "vithchataya.saharat@gmail.com",
//     "exp": 1721048992,
//     "iss": "1"
// }

type ClaimsData struct {
	Email    string `json:"email"`
	ExpireAt string `json:"exp"`
	UserID   string `json:"iss"`
}

type AuthService interface {
	SignIn(credential SignInRequest) (*TokenData, error)
	SignUp(credential SignUpRequest) (*TokenData, error)
	VerityToken(access_token string) (*ClaimsData, error)
}
