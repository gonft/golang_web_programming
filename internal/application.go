package internal

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	membership, err := app.repository.Create(request)
	return CreateResponse{membership.UserName, membership.MembershipType}, err
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	return UpdateResponse{"1", "jenny", "naver"}, nil
}

func (app *Application) Delete(id string) error {
	return nil
}
