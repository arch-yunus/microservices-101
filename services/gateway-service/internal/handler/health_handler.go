package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

// HealthHandler sistemin saglık durumunu doner
func HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "UP",
		"service": "Gateway-Service",
		"version": "1.0.1",
	})
}
