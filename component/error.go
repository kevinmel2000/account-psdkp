package component

import (
"math/rand"
"net/http"
"strconv"

log "github.com/sirupsen/logrus"
"github.com/google/jsonapi"
"github.com/labstack/echo"
"time"
)

const (
	CHARS = "abcdefghijklmnopqrstuvwxyz0123456789"
)

// Error setting error data
type Error struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Type   string `json:"type,omitempty"`
}

// ResponseErrorDocument setting error document
type ResponseErrorDocument struct {
	Errors []Error `json:"errors"`
	Code   int     `json:"-"`
}

// RandomString function for creating randomized string
// length int to set length of string
func randomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	charsLength := len(CHARS)
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = CHARS[rand.Intn(charsLength)]
	}
	return string(result)
}

// Error function for setting error detail
func (e *ResponseErrorDocument) Error() string {
	return e.Errors[0].Detail
}

// NewAppError function for serving application error to json
// code int http code
// message string error message
// eType error type whether direct discount or voucher
func NewAppError(code int, message string, eType ...string) *ResponseErrorDocument {
	e := &ResponseErrorDocument{}
	err := Error{
		ID:     randomString(10),
		Status: strconv.Itoa(code),
		Title:  "error",
		Detail: message,
	}
	// check error type
	if len(eType) > 0 {
		err.Type = eType[0]
	}

	e.Errors = append(e.Errors, err)
	e.Code = code

	return e
}

// AppHTTPErrorHandler variable to ser error handler
var AppHTTPErrorHandler = func(err error, c echo.Context) {
	if he, ok := err.(*ResponseErrorDocument); ok {
		if !c.Response().Committed {
			c.Response().Header().Set(echo.HeaderContentType, jsonapi.MediaType)
			c.Response().WriteHeader(he.Code)
			c.JSON(he.Code, he)
		}
		log.Error(he)
		return
	}

	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Error()
	}

	if !c.Response().Committed {
		c.Response().Header().Set(echo.HeaderContentType, jsonapi.MediaType)
		c.Response().WriteHeader(code)
		c.JSON(code, NewAppError(code, msg))
	}
	log.Error(err)
}

