package lib

import (
	middleware "date-rate/app/middelware"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tokens"
)

func setAuthToken(app core.App, c echo.Context, user *models.Record) error {
	if !user.Verified() {
		return nil
	}

	s, tokenErr := tokens.NewRecordAuthToken(app, user)
	if tokenErr != nil {
		return fmt.Errorf("login failed")
	}

	SetCookie(c, s)

	return nil
}

func SetCookie(c echo.Context, token string) {
	if token == "" {
		c.SetCookie(&http.Cookie{
			Name:     middleware.AuthCookieName,
			Value:    "",
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			MaxAge:   -1,
		})
		return
	}

	c.SetCookie(&http.Cookie{
		Name:     middleware.AuthCookieName,
		Value:    token,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	})
}

func Login(e *core.ServeEvent, c echo.Context, email string, password string) error {
	user, err := e.App.Dao().FindAuthRecordByEmail("users", email)
	if err != nil {
		return fmt.Errorf("Login failed")
	}

	valid := user.ValidatePassword(password)
	if !valid {
		return fmt.Errorf("Login failed")
	}

	return setAuthToken(e.App, c, user)
}

func Register(e *core.ServeEvent, c echo.Context, email string, username string, password string) error {
	collection, err := e.App.Dao().Clone().FindCollectionByNameOrId("users")
	if err != nil {
		return err
	}

	newUser := models.NewRecord(collection)
	if err := newUser.SetEmail(email); err != nil {
		return err
	}
	if err := newUser.SetUsername(username); err != nil {
		return err
	}
	if err := newUser.SetPassword(password); err != nil {
		return err
	}
	if err = e.App.Dao().SaveRecord(newUser); err != nil {
		return err
	}

	return setAuthToken(e.App, c, newUser)
}
