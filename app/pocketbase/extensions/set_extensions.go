package extensions

import "github.com/pocketbase/pocketbase/core"

func SetPocketbaseExtensions(e *core.ServeEvent, app core.App) {
	googleFirstTimeLoginSetAvatar(e, app)
	googleFirstTimeLoginSetUsername(e, app)
	googleFirstTimeLoginSetOauth(e, app)
}
