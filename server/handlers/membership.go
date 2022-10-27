package handlers

import (
	"github.com/labstack/echo/v4"
	"golang_web_programming/server/services"
)

type MembershipHandler struct {
	services.MembershipService
}

func (m *MembershipHandler) Get(c echo.Context) error {
	return nil
}

func (m *MembershipHandler) GetByID(c echo.Context) error {
	return nil
}

func (m *MembershipHandler) Create(c echo.Context) error {
	return nil
}

func (m *MembershipHandler) Update(c echo.Context) error {
	return nil
}

func (m *MembershipHandler) Delete(c echo.Context) error {
	return nil
}
