package services

import (
	"golang_web_programming/server/model/dto"
	"golang_web_programming/server/model/vo"
	"golang_web_programming/server/repositories"
)

type MembershipService interface {
	Create(request dto.CreateRequest) (*vo.CreateResponse, error)
	Update(request dto.UpdateRequest) (*vo.UpdateResponse, error)
	Delete(id string) error
	GetAll() []vo.GetResponse
	GetByID(id string) (*vo.GetResponse, error)
}

type MembershipServiceContext struct {
	repo *repositories.MembershipRepository
}

func (m *MembershipServiceContext) Create(request dto.CreateRequest) (*vo.CreateResponse, error) {
	membership, err := m.repo.Create(request)
	if err != nil {
		return nil, err
	}
	return &vo.CreateResponse{ID: membership.UserName, MembershipType: membership.MembershipType}, nil
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

func (m *MembershipServiceContext) GetAll() []vo.GetResponse {
	memberships := m.repo.GetAll()
	var response []vo.GetResponse
	for _, membership := range memberships {
		response = append(response, vo.GetResponse{ID: membership.ID, UserName: membership.UserName, MembershipType: membership.MembershipType})
	}
	return response
}

func (m *MembershipServiceContext) GetByID(id string) (*vo.GetResponse, error) {
	membership, err := m.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &vo.GetResponse{ID: membership.ID, UserName: membership.UserName, MembershipType: membership.MembershipType}, nil
}
