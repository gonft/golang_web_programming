package repositories

import (
	"errors"
	"golang.org/x/exp/slices"
	"golang_web_programming/server/model"
	"golang_web_programming/server/model/dto"
	"strconv"
)

var UserNameAlreadyExistsError = errors.New("이미 등록된 사용자 이름이 존재합니다.")
var UserIdNotFoundError = errors.New("사용자 이름을 찾을 수 없습니다.")
var UserNameEmptyError = errors.New("사용자 이름이 비어있습니다.")
var MembershipTypeEmptyError = errors.New("멤버십 타입이 비어있습니다.")
var MembershipTypeInvalidError = errors.New("멤버십 타입이 잘못되었습니다.")

type MembershipRepository struct {
	data map[string]model.Membership
}

// exists `id`를 가진 멤버쉽이 존재하는지 확인한다.
func (r *MembershipRepository) exists(id string) bool {
	_, ok := r.data[id]
	return ok
}

// existsName `name`을 가진 멤버쉽이 존재하는지 확인한다.
func (r *MembershipRepository) existsName(name string) bool {
	for _, value := range r.data {
		if value.UserName == name {
			return true
		}
	}
	return false
}

// Create 멤버쉽을 생성한다.
func (r *MembershipRepository) Create(request dto.CreateRequest) (*model.Membership, error) {
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
	membership := model.Membership{ID: strconv.Itoa(len(r.data) + 1), UserName: request.UserName, MembershipType: request.MembershipType}
	r.data[membership.ID] = membership
	return &membership, nil
}

// Update 멤버쉽을 업데이트한다.
func (r *MembershipRepository) Update(request dto.UpdateRequest) (*model.Membership, error) {
	// 사용자 ID를 찾을수 없는경우 에러
	if !r.exists(request.ID) {
		return nil, UserIdNotFoundError
	}
	// 먼저 멤버쉽을 찾는다
	membership := r.data[request.ID]
	switch {
	// 변경하려는 멤버쉽의 유저 이름이 기존과 이름과 서로 다르고 해당하는 이름이 이미 존재하는경우 에러
	case membership.UserName != request.UserName && r.existsName(request.UserName):
		return nil, UserNameAlreadyExistsError
	case request.UserName == "":
		return nil, UserNameEmptyError
	case request.MembershipType == "":
		return nil, MembershipTypeEmptyError
	case !slices.Contains([]string{"naver", "toss", "payco"}, request.MembershipType):
		return nil, MembershipTypeInvalidError
	}

	membership.UserName = request.UserName
	membership.MembershipType = request.MembershipType
	r.data[request.UserName] = membership
	return &membership, nil
}

// Delete 멤버쉽을 삭제한다.
func (r *MembershipRepository) Delete(id string) error {
	switch {
	case id == "" || !r.exists(id):
		return UserIdNotFoundError
	}
	delete(r.data, id)
	return nil
}

// GetAll 멤버쉽을 모두 가져온다.
func (r *MembershipRepository) GetAll() []model.Membership {
	var memberships []model.Membership
	for _, value := range r.data {
		memberships = append(memberships, value)
	}
	return memberships
}

// GetByID `id`를 가진 멤버쉽을 가져온다.
func (r *MembershipRepository) GetByID(id string) (*model.Membership, error) {
	if !r.exists(id) {
		return nil, UserIdNotFoundError
	}
	membership := r.data[id]
	return &membership, nil
}

// NewRepository 멤버쉽 리포지토리를 생성한다.
func NewRepository(data map[string]model.Membership) *MembershipRepository {
	return &MembershipRepository{data: data}
}
