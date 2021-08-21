// +build !headless

package systray

import (
	"fmt"

	"github.com/getlantern/i18n"
	"deng/src/github.com/getlantern/systray"

	"deng/src/github.com/getlantern/flashlight/app"
	"deng/src/github.com/getlantern/flashlight/icons"
	"deng/src/github.com/getlantern/flashlight/ui"
)

func RunOnSystrayReady(f func()) {
	systray.Run(f)
}

func QuitSystray() {
	systray.Quit()
}

func ConfigureSystemTray(a *app.App) error {
	icon, err := icons.Asset("icons/16on.ico")
	if err != nil {
		return fmt.Errorf("Unable to load icon for system tray: %v", err)
	}
	systray.SetIcon(icon)
	systray.SetTooltip("Lantern")
	show := systray.AddMenuItem(i18n.T("TRAY_SHOW_LANTERN"), i18n.T("SHOW"))
	quit := systray.AddMenuItem(i18n.T("TRAY_QUIT"), i18n.T("QUIT"))
	go func() {
		for {
			select {
			case <-show.ClickedCh:
				ui.Show()
			case <-quit.ClickedCh:
				a.Exit(nil)
				return
			}
		}
	}()

	return nil
}
