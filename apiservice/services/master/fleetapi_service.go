package master

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Master) GetHome(context echo.Context) error {
	return context.JSON(http.StatusOK, "Hello World")
}
