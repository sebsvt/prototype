package domain

type Membership struct {
	ID       int    `db:"id"`
	MemberID int    `db:"member_id"`
	StudioID int    `db:"studio_id"`
	Role     string `db:"role"`
}

type MembershipRepository interface {
	Add(Membership) error
	Find(member_id int, studio_id int) (*Membership, error)
	Remove(member_id int, studio_id int) error
	All(member_id int) ([]Membership, error)
}
