package extensions

import "github.com/pocketbase/pocketbase/core"

func googleFirstTimeLoginSetUsername(e *core.ServeEvent, app core.App) {
	e.App.OnRecordAfterAuthWithOAuth2Request("users").Add(func(e *core.RecordAuthWithOAuth2Event) error {
		if !e.IsNewRecord {
			return nil
		}

		user, err := app.Dao().FindAuthRecordByEmail("users", e.OAuth2User.Email)
		if err != nil {
			return err
		}

		if err := user.SetUsername(e.OAuth2User.Name); err != nil {
			return err
		}

		if err := app.Dao().SaveRecord(user); err != nil {
			return err
		}

		return nil
	})
}
