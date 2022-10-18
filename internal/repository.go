package internal

import (
	"errors"
	"golang.org/x/exp/slices"
	"strconv"
)

var UserNameAlreadyExistsError = errors.New("이미 등록된 사용자 이름이 존재합니다.")
var UserIdNotFoundError = errors.New("사용자 이름을 찾을 수 없습니다.")
var UserNameEmptyError = errors.New("사용자 이름이 비어있습니다.")
var MembershipTypeEmptyError = errors.New("멤버십 타입이 비어있습니다.")
var MembershipTypeInvalidError = errors.New("멤버십 타입이 잘못되었습니다.")

type Repository struct {
	data map[string]Membership
}

func (r *Repository) exists(id string) bool {
	_, ok := r.data[id]
	return ok
}

func (r *Repository) existsName(name string) bool {
	for _, value := range r.data {
		if value.UserName == name {
			return true
		}
	}
	return false
}

func (r *Repository) Create(request CreateRequest) (*Membership, error) {
	switch {
	case r.existsName(request.UserName):
		return nil, UserNameAlreadyExistsError
	case request.UserName == "":
		return nil, UserNameEmptyError
	case request.MembershipType == "":
		return nil, MembershipTypeEmptyError
	case !slices.Contains([]string{"naver", "toss", "payco"}, request.MembershipType):
		return nil, MembershipTypeInvalidError
	}
	membership := Membership{ID: strconv.Itoa(len(r.data) + 1), UserName: request.UserName, MembershipType: request.MembershipType}
	r.data[membership.ID] = membership
	return &membership, nil
}

func (r *Repository) Update(request UpdateRequest) (*Membership, error) {
	switch {
	// 사용자 ID를 찾을수 없는경우 에러
	case !r.exists(request.ID):
		return nil, UserIdNotFoundError
	case r.existsName(request.UserName):
		return nil, UserNameAlreadyExistsError
	case request.UserName == "":
		return nil, UserNameEmptyError
	case request.MembershipType == "":
		return nil, MembershipTypeEmptyError
	case !slices.Contains([]string{"naver", "toss", "payco"}, request.MembershipType):
		return nil, MembershipTypeInvalidError
	}
	membership := r.data[request.ID]
	membership.UserName = request.UserName
	membership.MembershipType = request.MembershipType
	r.data[request.UserName] = membership
	return &membership, nil
}

func (r *Repository) Delete(id string) error {
	switch {
	case id == "" || !r.exists(id):
		return UserIdNotFoundError
	}
	delete(r.data, id)
	return nil
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}
