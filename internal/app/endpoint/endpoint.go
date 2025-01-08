package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	Days() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {

	return &Endpoint{

		s: s,
	}

}

func (e *Endpoint) Status(c echo.Context) error {

	days := e.s.Days()

	start := fmt.Sprintf("Days Left:%d", days)

	err := c.String(http.StatusOK, start)

	if err != nil {
		return err
	}

	return nil
}
