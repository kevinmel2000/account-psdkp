package component

import (
	"encoding/json"
	"strings"
	"github.com/labstack/echo"
)

const AppJSONHeader = "application/vnd.api+json" //json header

//CustomBinder for checking header
type CustomBinder struct {}

//Bind CustomBinder for checking header
//i interface
//c context
func (cb *CustomBinder) Bind(i interface{},c echo.Context)(err error) {
	ct := c.Request().Header.Get(echo.HeaderContentType)

	err = echo.ErrUnsupportedMediaType
	if strings.HasPrefix(ct, AppJSONHeader){
		err = json.NewDecoder(c.Request().Body).Decode(i)

		return  err
	}

	return err
}
