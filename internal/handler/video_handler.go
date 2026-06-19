package handler

import (
	"StreamRoom/internal/service"
	"StreamRoom/internal/views"
	"fmt"
	"mime/multipart"

	"net/http"

	"github.com/labstack/echo/v4"
)

type VideoHandler struct {
	s *service.VideoService
	r *service.RoomService
}

func NewVideoHandler(group *echo.Group, videoService *service.VideoService, roomService *service.RoomService) {
	h := &VideoHandler{s: videoService, r: roomService}
	group.POST("video/upload", h.UploadVideo)
}

func (h *VideoHandler) UploadVideo(c echo.Context) error {
	file, err := c.FormFile("video")
	if err != nil {
		return err
	}
	var src multipart.File
	src, err = file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, views.Response{Message: fmt.Errorf("failed to open file :: %w", err).Error()})
	}
	defer src.Close()
	url, err := h.s.UploadToStorage(src, file.Filename)

	return c.JSON(http.StatusOK, "upload success :: "+url)
}
