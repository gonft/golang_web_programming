package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang_web_programming/server/middlewares"
	"golang_web_programming/server/services"
)

type Handlers struct {
	MembershipHandler
}

func New(s *services.Service) *Handlers {
	return &Handlers{
		MembershipHandler: MembershipHandler{s.MembershipService},
	}
}

func SetApi(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {
	g := e.Group("/api/v1")
	if m != nil {
		g.Use(m)
	}

	// Membership
	g.GET("/membership", h.MembershipHandler.GetAll)
	g.GET("/membership/:id", h.MembershipHandler.GetByID)
	g.POST("/membership", h.MembershipHandler.Create)
	g.PUT("/membership/:id", h.MembershipHandler.Update)
	g.DELETE("/membership/:id", h.MembershipHandler.Delete)
}

func Echo() *echo.Echo {
	e := echo.New()

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		req := c.Request()
		res := c.Response()

		var reqb map[string]interface{}
		err := json.Unmarshal(reqBody, &reqb)
		if err != nil {
			fmt.Println(err)
		}

		var resb map[string]interface{}
		err = json.Unmarshal(reqBody, &resb)
		if err != nil {
			fmt.Println(err)
		}

		reqMiddleware := middlewares.ReqRes{
			URI:          req.RequestURI,
			Method:       req.Method,
			RequestBody:  reqb,
			ResponseCode: res.Status,
			ResponseBody: resb,
		}
		s, _ := json.MarshalIndent(reqMiddleware, "", "  ")
		fmt.Println(string(s))
	}))
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	return e
}
