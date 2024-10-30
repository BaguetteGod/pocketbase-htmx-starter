package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/tools/filesystem"
)

var (
	MAX_FILE_SIZE    = 5000000
	VALID_FILE_TYPES = []string{"jpg", "png", "webp", "gif"}
	REGEX_USERNAME   = "^[a-z0-9]+$"
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

func GetDisplayName(c echo.Context) string {
	return apis.RequestInfo(c).AuthRecord.GetString("displayname")
}

func GetUsername(c echo.Context) string {
	return apis.RequestInfo(c).AuthRecord.GetString("username")
}

func GetAvatar(c echo.Context) string {
	user := apis.RequestInfo(c).AuthRecord
	filename := user.GetString("avatar")
	if filename == "" {
		return "https://upload.wikimedia.org/wikipedia/commons/7/7c/Profile_avatar_placeholder_large.png?20150327203541"
	}

	appUrl := GetAppUrl()
	avatar := fmt.Sprintf("%s/api/files/%s/%s/%s", appUrl, user.Collection().Id, user.Id, filename)
	return avatar
}

func TitleCase(s string) string {
	r := []rune(s)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}

func GetCurrentUrlString(c echo.Context) string {
	return c.Request().URL.String()
}

func GetAppUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("APP_URL")
}

func ValidateImage(file *filesystem.File) bool {
	stringSlice := strings.Split(file.Name, ".")
	fileType := stringSlice[len(stringSlice)-1]
	return file.Size <= int64(MAX_FILE_SIZE) && slices.Contains(VALID_FILE_TYPES, fileType)
}

type PocketBaseRequest struct {
	Data    map[string]interface{}
	Headers map[string]string
	Route   string
	Method  string
}

func NewPocketBaseRequest(pbr PocketBaseRequest) (*http.Response, error) {
	appUrl := GetAppUrl()

	url := fmt.Sprintf("%s/api/collections%s", appUrl, pbr.Route)

	jsonData, err := json.Marshal(pbr.Data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return (nil), (err)
	}

	req, err := http.NewRequest(pbr.Method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return (nil), (err)
	}

	if pbr.Method != "GET" {
		req.Header.Set("Content-type", "application/json")
	}
	for k := range maps.Keys(pbr.Headers) {
		req.Header.Set(k, pbr.Headers[k])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return (nil), (err)
	}
	defer resp.Body.Close()

	return (resp), (nil)
}
