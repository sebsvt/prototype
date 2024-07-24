package domain

type Profile struct {
	ProfileID   int    `db:"profile_id"`
	Avatar      string `db:"avatar"`
	FirstName   string `db:"firstname"`
	Surname     string `db:"surname"`
	Gender      string `db:"gender"` // male, female or prefer to not say
	Phone       string `db:"phone"`
	DateOfBirth string `db:"date_of_birth"`
	UserRef     int    `db:"user_ref"`
}

type ProfileRepository interface {
	Create(entity Profile) (int, error)
	FromUserRef(user_ref int) (*Profile, error)
	Update(entity Profile) error
}
