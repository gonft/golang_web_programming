package services

import (
	"golang_web_programming/server/repositories"
)

type Service struct {
	MembershipService
}

func New(repos *repositories.MembershipRepository) *Service {
	return &Service{
		MembershipService: &MembershipServiceContext{repos},
	}
}
