package routes

import "github.com/labstack/echo/v4"

type IRoutes interface {
	MasterRoutes(server *echo.Echo)
}
