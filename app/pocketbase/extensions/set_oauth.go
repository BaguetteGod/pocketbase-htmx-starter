package extensions

import "github.com/pocketbase/pocketbase/core"

func googleFirstTimeLoginSetOauth(e *core.ServeEvent, app core.App) {
	e.App.OnRecordAfterAuthWithOAuth2Request("users").Add(func(e *core.RecordAuthWithOAuth2Event) error {
		if !e.IsNewRecord {
			return nil
		}

		user, err := app.Dao().FindAuthRecordByEmail("users", e.OAuth2User.Email)
		if err != nil {
			return err
		}

		user.Set("oauth", true)

		if err := app.Dao().SaveRecord(user); err != nil {
			return err
		}

		return nil
	})
}
