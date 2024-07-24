package domain

import "time"

type Studio struct {
	StudioID    int        `db:"studio_id"`
	SubDomain   string     `db:"subdomain"`
	Picture     string     `db:"picture"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Address     string     `db:"address"`
	City        string     `db:"city"`
	ZipCode     string     `db:"zipcode"`
	State       string     `db:"state"`
	Country     string     `db:"country"`
	OwnerID     int        `db:"owner_id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

type StudioRepository interface {
	Create(entity Studio) (int, error)
	FromSubDomain(subdomain string) (*Studio, error)
	FromID(studio_id int) (*Studio, error)
	Update(entity Studio) error
}
