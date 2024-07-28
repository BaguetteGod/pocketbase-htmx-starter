package extensions

import (
	"context"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/tools/filesystem"
)

func googleFirstTimeLoginSetAvatar(e *core.ServeEvent, app core.App) {
	e.App.OnRecordAfterAuthWithOAuth2Request("users").Add(func(e *core.RecordAuthWithOAuth2Event) error {
		if !e.IsNewRecord {
			return nil
		}

		user, err := app.Dao().FindAuthRecordByEmail("users", e.OAuth2User.Email)
		if err != nil {
			return err
		}

		form := forms.NewRecordUpsert(app, user)
		if err := form.LoadData(map[string]any{
			"username": user.Username(),
			"email":    user.Email(),
		}); err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		file, err := filesystem.NewFileFromUrl(ctx, e.OAuth2User.AvatarUrl)
		if err != nil {
			app.Logger().Error("failed to fetch file from url", "error", err)
			return err
		}

		if err := form.AddFiles("avatar", file); err != nil {
			app.Logger().Error(err.Error())
			return err
		}

		if err := form.Submit(); err != nil {
			app.Logger().Error(err.Error())
			return err
		}

		return nil
	})
}
