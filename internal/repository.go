package internal

import (
	"errors"
	"golang.org/x/exp/slices"
)

var UserNameAlreadyExistsError = errors.New("이미 등록된 사용자 이름이 존재합니다.")
var UserNameEmptyError = errors.New("사용자 이름이 비어있습니다.")
var MembershipTypeEmptyError = errors.New("멤버십 타입이 비어있습니다.")
var MembershipTypeInvalidError = errors.New("멤버십 타입이 잘못되었습니다.")

type Repository struct {
	data map[string]Membership
}

func (r *Repository) exists(name string) bool {
	_, ok := r.data[name]
	return ok
}

func (r *Repository) Create(request CreateRequest) (*Membership, error) {
	switch {
	case r.exists(request.UserName):
		return nil, UserNameAlreadyExistsError
	case request.UserName == "":
		return nil, UserNameEmptyError
	case request.MembershipType == "":
		return nil, MembershipTypeEmptyError
	case !slices.Contains([]string{"naver", "toss", "payco"}, request.MembershipType):
		return nil, MembershipTypeInvalidError
	}
	membership := Membership{UserName: request.UserName, MembershipType: request.MembershipType}
	r.data[request.UserName] = membership
	return &membership, nil
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}
