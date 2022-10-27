package services

import (
	"golang_web_programming/server/model"
	"golang_web_programming/server/model/dto"
	"golang_web_programming/server/model/vo"
	"golang_web_programming/server/repositories"
)

type MembershipService interface {
	Create(request dto.CreateRequest) (*vo.CreateResponse, error)
	Update(request dto.UpdateRequest) (*vo.UpdateResponse, error)
	Delete(id string) error
	Get() ([]model.Membership, error)
	GetByID(id string) (model.Membership, error)
}

type MembershipServiceContext struct {
	repo *repositories.MembershipRepository
}

func (m *MembershipServiceContext) Create(request dto.CreateRequest) (*vo.CreateResponse, error) {
	membership, err := m.repo.Create(request)
	if err != nil {
		return nil, err
	}
	return &vo.CreateResponse{membership.UserName, membership.MembershipType}, nil
}

func (m *MembershipServiceContext) Update(request dto.UpdateRequest) (*vo.UpdateResponse, error) {
	membership, err := m.repo.Update(request)
	if err != nil {
		return nil, err
	}
	return &vo.UpdateResponse{membership.ID, membership.UserName, membership.MembershipType}, nil
}

func (m *MembershipServiceContext) Delete(id string) error {
	return m.repo.Delete(id)
}

func (m *MembershipServiceContext) Get() ([]model.Membership, error) {
	return []model.Membership{}, nil
}

func (m *MembershipServiceContext) GetByID(id string) (model.Membership, error) {
	return model.Membership{}, nil
}
