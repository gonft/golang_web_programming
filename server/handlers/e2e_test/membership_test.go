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

	t.Run("멤버쉽 생성", func(t *testing.T) {
		t.Run("네이버 멤버쉽 생성", func(t *testing.T) {
			// given 네이버 멤버쉽 생성
			e.POST("/api/v1/membership").
				WithJSON(
					dto.CreateRequest{
						MembershipType: "naver",
						UserName:       "jenny",
					},
				).
				Expect().
				Status(http.StatusCreated).
				JSON().Object().
				Value("id").String().NotEmpty()
			// when
			// then
		})
	})
}
