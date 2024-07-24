package services

import (
	"errors"
	"time"
)

var (
	ErrSubDomainAlreadyInUse = errors.New("studio's subdomain already in use")
)

type StudioResponse struct {
	StudioID    int        `json:"studio_id"`
	SubDomain   string     `json:"subdomain"`
	Picture     string     `json:"picture"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Address     string     `json:"address"`
	City        string     `json:"city"`
	ZipCode     string     `json:"zipcode"`
	State       string     `json:"state"`
	Country     string     `json:"country"`
	OwnerID     int        `json:"owner_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type StudioRequest struct {
	SubDomain   string `json:"subdomain"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ZipCode     string `json:"zipcode"`
	State       string `json:"state"`
	Country     string `json:"country"`
}

type StudioUpdateRequest struct {
	SubDomain   string `json:"subdomain"`
	Picture     string `json:"picture"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	City        string `json:"city"`
	ZipCode     string `json:"zipcode"`
	State       string `json:"state"`
	Country     string `json:"country"`
}

type MemberRequest struct {
	MemberID int    `json:"member_id"`
	StudioID int    `json:"studio_id"`
	Role     string `json:"role"`
}

type MemberRemoveRequest struct {
	MemberID int `json:"member_id"`
	StudioID int `json:"studio_id"`
}

type StudioService interface {
	CreateNewStudio(entity StudioRequest, owner_id int) (int, error)
	GetStudioBySubDomain(subdomain string) (*StudioResponse, error)
	GetAllStudiosFromUser(user_id int) ([]StudioResponse, error)
	UpdateStudio(entity StudioUpdateRequest) error
	AddNewMemberToStudioWithRole(entity MemberRequest) error
	RemoveMemberFromStudio(entity MemberRemoveRequest) error
	ChangeMemberRole(entity MemberRequest) error
}
