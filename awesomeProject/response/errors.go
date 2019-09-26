package response

import (
	"net/http"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"

	"github.com/labstack/echo"
)

type HandlerError struct {
	code    int      // http status code
	message string   //client user friendly error message
	err     string   //for log
	details []string // additional details
}

// Most used http responses
var (
	// ErrBadRequest -- code 400
	ErrBadRequest = NewError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), http.StatusText(http.StatusBadRequest))
	// ErrUnauthorized -- code 401
	ErrUnauthorized = NewError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized))
	// ErrForbidden -- code 403
	ErrForbidden = NewError(http.StatusForbidden, http.StatusText(http.StatusForbidden), http.StatusText(http.StatusForbidden))
	// ErrNotFound -- code 404
	ErrNotFound = NewError(http.StatusNotFound, http.StatusText(http.StatusNotFound), http.StatusText(http.StatusNotFound))
	// ErrInternalError -- code 500
	ErrInternalError = NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError))
)

func (e HandlerError) Error() string {
	return e.message
}

// SetDetails -- add Details to error response
func (e *HandlerError) SetDetails(details ...string) *HandlerError {
	err := *e
	err.details = append(err.details, details...)
	return &err
}

// SetMessage -- update error message from default value to custom value
func (e *HandlerError) SetMessage(msg string) *HandlerError {
	err := *e
	err.message = msg
	return &err
}

// SetCode -- update code from default value to custom value
func (e *HandlerError) SetCode(code int) *HandlerError {
	err := *e
	err.code = code
	return &err
}

// SetError -- update err from default value to custom error log message
func (e *HandlerError) SetError(logMessage string) *HandlerError {
	err := *e
	err.err = logMessage
	return &err
}

// NewError -- handler error constructor
func NewError(code int, msg, err string) *HandlerError {
	return &HandlerError{code: code, message: msg, err: err}
}

// HTTPErrorHandler -- default HTTP error handler. It sends a JSON response
// with status code.
func HTTPErrorHandler(err error, c echo.Context) {
	var (
		code    = http.StatusInternalServerError
		msg     interface{}
		details = []string{}
	)

	switch e := err.(type) {
	case *echo.HTTPError:
		code = e.Code
		msg = e.Message
	case *HandlerError:
		code = e.code
		msg = e.message
		details = e.details
	default:
		if c.Echo().Debug {
			msg = e.Error()
		} else {
			msg = http.StatusText(code)
		}
	}
	if _, ok := msg.(string); ok {
		msg = echo.Map{"code": code, "message": msg, "details": details}
	}

	if !c.Response().Committed {
		//even if error occurs - we do nothing, so just don't check it
		c.JSON(code, msg)
	}
}

// HandleModelError handle model error return status code
func HandleModelError(e error) *HandlerError {
	switch e {
	case errors.Cause(mgo.ErrNotFound):
		return ErrNotFound
	default:
		return ErrInternalError
	}
}
