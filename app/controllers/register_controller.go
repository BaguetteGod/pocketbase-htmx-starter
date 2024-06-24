package controllers

import (
	"bytes"
	"date-rate/app/forms"
	"date-rate/app/lib"
	"date-rate/app/views/register"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func handleVerificationEmailRequest(email string, error error) {
	if error != nil {
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := fmt.Sprintf("%s/api/collections/users/request-verification", os.Getenv("APP_URL"))

	bodyValue := map[string]interface{}{
		"email": email,
	}
	jsonData, err := json.Marshal(bodyValue)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
}

func RegisterRegisterRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("/register", func(c echo.Context) error {
		return lib.Render(c, register.RegisterPage(forms.RegisterFormValue{}, ""))
	})

	group.POST("/register", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, register.RegisterPage(form, string(formErrors)))
		}

		err = lib.Register(e, c, form.Email, form.Username, form.Password)
		handleVerificationEmailRequest(form.Email, err)

		return c.Redirect(302, "/register/success")
	})

	group.GET("/register/success", func(c echo.Context) error {
		return lib.Render(c, register.RegisterSuccessPage())
	})
}
