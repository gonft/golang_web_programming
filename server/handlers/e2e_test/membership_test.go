package e2e_test

import (
	"github.com/gavv/httpexpect/v2"
	"golang_web_programming/server"
	"golang_web_programming/server/handlers"
	"golang_web_programming/server/model/dto"
	"net/http"
	"testing"
)

func TestMembership(t *testing.T) {
	s := server.NewDefaultServer()
	echoServer := handlers.Echo()
	handlers.SetApi(echoServer, s.Handlers, nil)

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(echoServer),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	t.Run("ATDD", func(t *testing.T) {
		t.Run("멤버쉽 생성 삭재후 재삭제 불가능", func(t *testing.T) {
			// given: 멤버쉽을 생성 한다
			e.POST("/api/v1/membership").
				WithJSON(
					dto.CreateRequest{
						MembershipType: "naver",
						UserName:       "jenny",
					},
				).
				Expect().
				Status(http.StatusCreated)
			// given: 멤버쉽을 삭제 한다
			e.DELETE("/api/v1/membership/1").
				Expect().
				Status(http.StatusOK)
			// given: 멤버십을 다시 삭제할 수 없다
			e.DELETE("/api/v1/membership/1").
				Expect().
				Status(http.StatusBadRequest)
		})
	})

}
