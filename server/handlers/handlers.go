package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang_web_programming/server/services"
	"strings"
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
	g.Use(m)

	// Membership
	g.GET("/membership", h.MembershipHandler.GetAll)
	g.GET("/membership/:id", h.MembershipHandler.GetByID)
	g.POST("/membership", h.MembershipHandler.Create)
	g.PUT("/membership/:id", h.MembershipHandler.Update)
	g.DELETE("/membership/:id", h.MembershipHandler.Delete)
}

func Echo() *echo.Echo {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	return e
}
