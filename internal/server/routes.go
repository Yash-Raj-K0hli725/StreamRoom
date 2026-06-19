package server

import (
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterRoutes() http.Handler {

	/*--------prefix---------*/
	apiV1Group := s.e.Group("/api/v1")

	apiV1Group.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}))
	//apiV1Group.Use(handlers.AuthMiddleware)

	/*-------------public group---------------------*/
	publicGroup := s.e.Group("/public")

	/*-------------Service Layer------------*/

	/*-------------Handler Layer-------------*/
	//##-with auth-##

	//##-without auth-##
	publicGroup.GET("/health", s.healthHandler)

	return s.e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "good")
}
