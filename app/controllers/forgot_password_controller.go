package controllers

import (
	"bytes"
	"date-rate/app/forms"
	"date-rate/app/lib"
	"date-rate/app/views/forgot_password"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func handleForgotPasswordRequest(form forms.ForgotPasswordFormValue) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := fmt.Sprintf("%s/api/collections/users/request-password-reset", os.Getenv("APP_URL"))

	jsonData, err := json.Marshal(form)
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

func handleConfirmPasswordResetRequest(form forms.ConfirmPasswordResetFormValue) *http.Response {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := fmt.Sprintf("%s/api/collections/users/confirm-password-reset", os.Getenv("APP_URL"))

	jsonData, err := json.Marshal(form)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.Status != "204" {
		return resp
	}

	return nil
}

func RegisterForgotPasswordRoutes(e *core.ServeEvent, group echo.Group) {
	group.GET("/forgot-password", func(c echo.Context) error {
		return lib.Render(c, forgot_password.ForgotPasswordPage(forms.ForgotPasswordFormValue{}, ""))
	})

	group.POST("/forgot-password", func(c echo.Context) error {
		form := forms.GetForgotPasswordFormValue(c)
		err := form.Validate()
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, forgot_password.ForgotPasswordPage(forms.ForgotPasswordFormValue{}, string(formErrors)))
		}

		handleForgotPasswordRequest(form)

		return c.Redirect(302, "/forgot-password/success")
	})

	group.GET("/forgot-password/success", func(c echo.Context) error {
		return lib.Render(c, forgot_password.ForgotPasswordSuccessPage())
	})

	group.GET("/confirm-password-reset/:token", func(c echo.Context) error {
		token := c.PathParam("token")
		form := forms.ConfirmPasswordResetFormValue{Token: token}

		return lib.Render(c, forgot_password.ConfirmPasswordResetPage(form, ""))
	})

	group.POST("/confirm-password-reset/:token", func(c echo.Context) error {
		form := forms.GetConfirmPasswordResetForm(c)
		err := form.Validate()
		if err != nil {
			formErrors, _ := json.Marshal(err)
			return lib.Render(c, forgot_password.ConfirmPasswordResetPage(form, string(formErrors)))
		}

		resp := handleConfirmPasswordResetRequest(form)
		if resp != nil {
			errorMap := map[string]interface{}{
				"token": "Link is invalid or has expired",
			}
			formErrors, _ := json.Marshal(errorMap)
			return lib.Render(c, forgot_password.ConfirmPasswordResetPage(form, string(formErrors)))
		}

		return c.Redirect(302, "/confirm-password-reset/success")
	})

	group.GET("/confirm-password-reset/success", func(c echo.Context) error {
		return lib.Render(c, forgot_password.ConfirmPasswordSuccessPage())
	})
}
