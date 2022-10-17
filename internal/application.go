package internal

import "errors"

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

var UserNameAlreadyExistsError = errors.New("이미 등록된 사용자 이름이 존재합니다.")

func (app *Application) Create(request CreateRequest) (*CreateResponse, error) {
	return app.repository.Save(request)
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	return UpdateResponse{}, nil
}

func (app *Application) Delete(id string) error {
	return nil
}
