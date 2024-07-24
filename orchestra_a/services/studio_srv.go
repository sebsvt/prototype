package services

import (
	"database/sql"
	"time"

	"github.com/sebsvt/prototype/orchestra/domain"
	"github.com/sebsvt/prototype/orchestra/logs"
)

type studioService struct {
	studio_repo domain.StudioRepository
	member_repo domain.MembershipRepository
}

func NewStudioService(studio_repo domain.StudioRepository, member_repo domain.MembershipRepository) StudioService {
	return studioService{studio_repo: studio_repo, member_repo: member_repo}
}

// CreateNewStudio implements StudioService.
func (srv studioService) CreateNewStudio(entity StudioRequest, owner_id int) (int, error) {
	// is it subdomain already in use
	studio, err := srv.studio_repo.FromSubDomain(entity.SubDomain)
	if err != nil && err != sql.ErrNoRows {
		logs.Error(err)
		return 0, err
	}
	if studio != nil {
		logs.Error(ErrSubDomainAlreadyInUse)
		return 0, ErrSubDomainAlreadyInUse
	}
	// create new subdomain
	new_studio := domain.Studio{
		SubDomain:   entity.SubDomain,
		Picture:     "",
		Name:        entity.Name,
		Description: entity.Description,
		Address:     entity.Address,
		City:        entity.City,
		ZipCode:     entity.ZipCode,
		State:       entity.State,
		Country:     entity.Country,
		OwnerID:     owner_id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	}
	studio_id, err := srv.studio_repo.Create(new_studio)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	// add owner to new member
	if err := srv.AddNewMemberToStudioWithRole(MemberRequest{
		MemberID: owner_id,
		StudioID: studio_id,
		Role:     "owner",
	}); err != nil {
		logs.Error(err)
		return 0, err
	}
	return studio_id, nil
}

// GetStudioBySubDomain implements StudioService.
func (srv studioService) GetStudioBySubDomain(subdomain string) (*StudioResponse, error) {
	entity, err := srv.studio_repo.FromSubDomain(subdomain)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &StudioResponse{
		StudioID:    entity.StudioID,
		SubDomain:   entity.SubDomain,
		Picture:     entity.Picture,
		Name:        entity.Name,
		Description: entity.Description,
		Address:     entity.Address,
		City:        entity.City,
		ZipCode:     entity.ZipCode,
		State:       entity.State,
		Country:     entity.Country,
		OwnerID:     entity.OwnerID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}, nil
}

// AddNewMemberToStudioWithRole implements StudioService.
func (srv studioService) AddNewMemberToStudioWithRole(entity MemberRequest) error {
	new_member := domain.Membership{
		MemberID: entity.MemberID,
		StudioID: entity.StudioID,
		Role:     entity.Role,
	}
	// if err is nil it will return nil, we don't have to have a condition to check error is nil then
	if err := srv.member_repo.Add(new_member); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// ChangeMemberRole implements StudioService.
func (srv studioService) ChangeMemberRole(entity MemberRequest) error {
	return nil
}

// RemoveMemberFromStudio implements StudioService.
func (srv studioService) RemoveMemberFromStudio(entity MemberRemoveRequest) error {
	if err := srv.member_repo.Remove(entity.MemberID, entity.StudioID); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// UpdateStudio implements StudioService.
func (srv studioService) UpdateStudio(entity StudioUpdateRequest) error {
	updated_studio := domain.Studio{
		SubDomain:   entity.SubDomain,
		Picture:     entity.Picture,
		Name:        entity.Name,
		Description: entity.Description,
		Address:     entity.Address,
		City:        entity.City,
		ZipCode:     entity.ZipCode,
		State:       entity.State,
		Country:     entity.Country,
		UpdatedAt:   time.Now(),
	}
	err := srv.studio_repo.Update(updated_studio)
	return err
}

func (srv studioService) GetAllStudiosFromUser(user_id int) ([]StudioResponse, error) {
	var studios []StudioResponse
	member_lists, err := srv.member_repo.All(user_id)
	if err != nil {
		return nil, err
	}
	for _, member := range member_lists {
		studio, err := srv.studio_repo.FromID(member.StudioID)
		if err != nil {
			return nil, err
		}
		studios = append(studios, StudioResponse{
			StudioID:    studio.StudioID,
			SubDomain:   studio.SubDomain,
			Picture:     studio.Picture,
			Name:        studio.Name,
			Description: studio.Description,
			Address:     studio.Address,
			City:        studio.City,
			ZipCode:     studio.ZipCode,
			State:       studio.State,
			Country:     studio.Country,
			OwnerID:     studio.OwnerID,
			CreatedAt:   studio.CreatedAt,
			UpdatedAt:   studio.UpdatedAt,
			DeletedAt:   studio.DeletedAt,
		})
	}
	return studios, nil
}
