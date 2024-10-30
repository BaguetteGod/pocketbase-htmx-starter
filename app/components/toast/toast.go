package toast

import (
	"cmp"
	"encoding/json"

	"github.com/labstack/echo/v5"
)

const (
	INFO    = "info"
	SUCCESS = "success"
	WARNING = "warning"
	DANGER  = "danger"
)

type Toast struct {
	Level   string `json:"level"`
	Title   string `json:"title"`
	Message string `json:"message"`
	Open    bool   `json:"open"`
}

func (t Toast) jsonify() (string, error) {
	eventMap := map[string]Toast{}
	eventMap["makeToast"] = t
	jsonData, err := json.Marshal(eventMap)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func New(c echo.Context, toast Toast) {
	toast.Level = cmp.Or(toast.Level, SUCCESS)
	toast.Open = cmp.Or(toast.Open, false)
	toast.Title = cmp.Or(toast.Title, "Your changes have been saved!")
	toast.SetHXTriggerHeader(c)
}

func (t Toast) SetHXTriggerHeader(c echo.Context) {
	jsonData, _ := t.jsonify()
	c.Response().Header().Set("HX-Trigger", jsonData)
}
