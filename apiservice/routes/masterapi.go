package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/selector/roundrobin"
	mast "tionyxtrack/apiservice/services/master"
)

var masterService micro.Service

func init() {
	masterService = micro.NewService(micro.Name("go.micro.srv.master"), micro.Selector(roundrobin.NewSelector()))
	masterService.Init()

}
func New() IRoutes {
	return Routes{}
}
func (Routes) MasterRoutes(server *echo.Echo) {
	routes := mast.New(masterService)
	server.GET("/", routes.GetHome)
}
