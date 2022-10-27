package server

import (
	"fmt"
	"golang_web_programming/server/handlers"
	"golang_web_programming/server/model"
	"golang_web_programming/server/repositories"
	"golang_web_programming/server/services"
	"log"
)

const _defaultPort = 8080

type Server struct {
	Handlers *handlers.Handlers
}

func NewDefaultServer() *Server {
	// 메모리 데이터 (유사 DB) 생성
	data := map[string]model.Membership{}
	// 서비스 생성
	service := services.New(repositories.NewRepository(data))
	// 핸들러 생성
	handler := handlers.New(service)
	return &Server{
		Handlers: handler,
	}
}

func (s *Server) Run() {
	e := handlers.Echo()
	handlers.SetApi(e, s.Handlers, nil)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}
