package internal

import "strconv"

type Repository struct {
	data map[string]Membership
}

func (r *Repository) Exists(name string) bool {
	_, ok := r.data[name]
	return ok
}

func (r *Repository) Save(request CreateRequest) (*CreateResponse, error) {
	if r.Exists(request.UserName) {
		return nil, UserNameAlreadyExistsError
	}
	r.data[request.UserName] = Membership{UserName: request.UserName, MembershipType: request.MembershipType}
	return &CreateResponse{strconv.Itoa(len(r.data)), request.MembershipType}, nil
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}
