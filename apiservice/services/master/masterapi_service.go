package master

import (
	"github.com/labstack/echo/v4"
	"tionyxtrack/apiservice/contracts/request"
)

func (m *Master) CreateCompany(context echo.Context) error {
	context.Get("ffd")
	return nil
}

func (m *Master) UpdateCompany(context echo.Context) error {

	return nil
}

func (m *Master) DeleteCompany(context echo.Context) error {

	return nil
}
func (m *Master) CreateContact(context echo.Context) error {
	u := new(request.ContactRequest)
	if err := context.Bind(u); err != nil {
		return err
	}
	return nil
}
