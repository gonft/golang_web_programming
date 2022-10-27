package server

import (
	"golang_web_programming/server/model/dto"
	"golang_web_programming/server/model/vo"
	"golang_web_programming/server/repositories"
)

type Application struct {
	repository repositories.MembershipRepository
}

func NewApplication(repository repositories.MembershipRepository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request dto.CreateRequest) (*vo.CreateResponse, error) {
	membership, err := app.repository.Create(request)
	if err != nil {
		return nil, err
	}
	return &vo.CreateResponse{membership.UserName, membership.MembershipType}, nil
}

func (app *Application) Update(request dto.UpdateRequest) (*vo.UpdateResponse, error) {
	membership, err := app.repository.Update(request)
	if err != nil {
		return nil, err
	}
	return &vo.UpdateResponse{membership.ID, membership.UserName, membership.MembershipType}, nil
}

func (app *Application) Delete(id string) error {
	return app.repository.Delete(id)
}
