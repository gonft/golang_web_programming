package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		// given: 어플리케이션 멤버쉽 레포에 사용자 jenny가 존재한다.
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		_, _ = app.Create(req)
		// when: 다시 jenny를 등록하려고 한다.
		res, err := app.Create(req)
		// then: 실패한다.
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		// given: 어플리케이션 멤버쉽 레포가 주어지고
		app := NewApplication(*NewRepository(map[string]Membership{}))
		// when: 이름 없이 사용자를 등록하려고 한다.
		req := CreateRequest{"", "naver"}
		res, err := app.Create(req)
		// then: 실패한다.
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		// given: 어플리케이션 멤버쉽 레포가 주어지고
		app := NewApplication(*NewRepository(map[string]Membership{}))
		// when: 이름 없이 사용자를 등록하려고 한다.
		req := CreateRequest{"jenny", ""}
		res, err := app.Create(req)
		// then: 실패한다.
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {

	})
}

func TestUpdate(t *testing.T) {
	t.Run("멤버십 정보를 갱신한다.", func(t *testing.T) {

	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {

	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {

	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {

	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {

	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {

	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {

	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {

	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {

	})
}