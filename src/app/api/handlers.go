package apiserver

import (
	"net/http"
	"github.com/labstack/echo/v4"
	
)


// GET /health-ping
func (s *apiServer) healthPing(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// insert additional handlers here