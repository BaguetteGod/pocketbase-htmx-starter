package lib

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

// Helper function to convert query parameter to integer
func ConvertToInt(param string, paramName string) (int, error) {
	val, err := strconv.Atoi(param)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Invalid "+paramName+" parameter")
	}
	return val, nil
}
