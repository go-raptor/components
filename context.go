package components

import (
	"errors"
	"net/http"

	"github.com/go-raptor/errs"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	Controller string
	Action     string
}

func (c *Context) JSONResponse(data interface{}, status ...int) error {
	if len(status) == 0 {
		status = append(status, http.StatusOK)
	}
	return c.Context.JSON(status[0], data)
}

func (c *Context) JSONError(err error, status ...int) error {
	var e *errs.Error
	if errors.As(err, &e) {
		return c.JSON(e.Code, e)
	}

	if len(status) == 0 {
		status = append(status, http.StatusInternalServerError)
	}
	return c.JSON(status[0], errs.NewError(status[0], err.Error()))
}
