package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pb-starter/app/components/inputs"
	"pb-starter/app/forms"
	"pb-starter/app/lib"
	"pb-starter/app/views/register"

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

		err = lib.Register(e, c, form.Email, form.Username, form.DisplayName, form.Password)
		handleVerificationEmailRequest(form.Email, err)

		return lib.Render(c, register.RegisterSuccessPage(form))
	})

	group.POST("/register/email", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{Name: "email", Label: "Email address", Value: form.Email, HxPost: "/register/email", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Text{Name: "email", Label: "Email address", Value: form.Email, HxPost: "/register/email"}.Comp())
	})

	group.POST("/register/username", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{Name: "username", Value: form.Username, HxPost: "/register/username", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Text{Name: "username", Value: form.Username, HxPost: "/register/username"}.Comp())
	})

	group.POST("/register/display-name", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Text{Name: "displayname", Label: "Display name", Value: form.DisplayName, HxPost: "/register/display-name", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Text{Name: "displayname", Label: "Display name", Value: form.DisplayName, HxPost: "/register/display-name"}.Comp())
	})

	group.POST("/register/password", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Password{Name: "password", Value: form.Password, HxPost: "/register/password", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Password{Name: "password", Value: form.Password, HxPost: "/register/password"}.Comp())
	})

	group.POST("/register/password-confirm", func(c echo.Context) error {
		form := forms.GetRegisterFormValue(c)
		err := form.Validate(e)
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, inputs.Password{Name: "passwordConfirm", Label: "Confirm password", Value: form.PasswordConfirm, HxPost: "/register/password-confirm", Error: string(formErrors)}.Comp())
		}

		return lib.Render(c, inputs.Password{Name: "passwordConfirm", Label: "Confirm password", Value: form.PasswordConfirm, HxPost: "/register/password-confirm"}.Comp())
	})
}
