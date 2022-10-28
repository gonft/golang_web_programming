package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang_web_programming/server/model/dto"
	"golang_web_programming/server/services"
	"golang_web_programming/server/utils"
	"net/http"
)

type MembershipHandler struct {
	services.MembershipService
}

func (m *MembershipHandler) GetAll(c echo.Context) error {
	res := m.MembershipService.GetAll()
	return c.JSON(http.StatusOK, res)
}

func (m *MembershipHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	res, err := m.MembershipService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (m *MembershipHandler) Create(c echo.Context) error {
	var req dto.CreateRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}
	_, err = m.MembershipService.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, "OK")
}

func (m *MembershipHandler) Update(c echo.Context) error {
	var req dto.UpdateRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}
	_, err = m.MembershipService.Update(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, "OK")
}

func (m *MembershipHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := m.MembershipService.Delete(id)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, "OK")
}
