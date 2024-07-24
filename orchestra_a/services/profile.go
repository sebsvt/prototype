package services

type ProfileRequest struct {
	FirstName   string `json:"firstname"`
	Surname     string `json:"surname"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
}

type ProfileResponse struct {
	UserRef     int    `json:"user_ref"`
	FirstName   string `json:"firstname"`
	Surname     string `json:"surname"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
}

type ProfileService interface {
	CreateNewProfile(entity ProfileRequest, user_ref int) error
	GetProfileFromUserID(user_id int) (*ProfileResponse, error)
}
