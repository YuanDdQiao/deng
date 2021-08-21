// +build headless

package systray

import (
	"deng/src/github.com/getlantern/flashlight/app"
)

func RunOnSystrayReady(f func()) {
	f()
}

func QuitSystray() {
}

func ConfigureSystemTray(a *app.App) error {
	return nil
}
