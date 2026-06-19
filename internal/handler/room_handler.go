package handler

import (
	"StreamRoom/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoomsHandler struct {
	s *service.RoomService
}

func NewRoomsHandler(group *echo.Group, service *service.RoomService) *RoomsHandler {
	h := &RoomsHandler{
		s: service,
	}
	group.POST("/rooms/create", h.CreateNewRoom)
	return h
}

func (h *RoomsHandler) CreateNewRoom(c echo.Context) error {
	response := h.s.CreateRoom("")
	return c.JSON(http.StatusOK, response)
}
