package internal

import (
	"errors"
	"strconv"
)

var UserNameAlreadyExistsError = errors.New("이미 등록된 사용자 이름이 존재합니다.")
var UserNameEmptyError = errors.New("사용자 이름이 비어있습니다.")
var MembershipTypeEmptyError = errors.New("멤버십 타입이 비어있습니다.")

type Repository struct {
	data map[string]Membership
}

func (r *Repository) Exists(name string) bool {
	_, ok := r.data[name]
	return ok
}

func (r *Repository) Save(request CreateRequest) (*CreateResponse, error) {
	switch {
	case r.Exists(request.UserName):
		return nil, UserNameAlreadyExistsError
	case request.UserName == "":
		return nil, UserNameEmptyError
	case request.MembershipType == "":
		return nil, MembershipTypeEmptyError
	}
	r.data[request.UserName] = Membership{UserName: request.UserName, MembershipType: request.MembershipType}
	return &CreateResponse{strconv.Itoa(len(r.data)), request.MembershipType}, nil
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}
