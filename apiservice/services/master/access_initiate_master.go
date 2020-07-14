package master

import (
	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro/v2"
)

type Master struct{}

type IMaster interface {
	GetHome(context echo.Context) error
	CreateCompany(context echo.Context) error
	UpdateCompany(context echo.Context) error
	DeleteCompany(context echo.Context) error
	CreateContact(context echo.Context) error
}

var mic micro.Service

func New(service micro.Service) IMaster {
	mic = service
	return &Master{}
}
