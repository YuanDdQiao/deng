// Package launcher configures Lantern to run on system start
package launcher

import (
	"fmt"
	"deng/src/github.com/kardianos/osext"
	"deng/src/github.com/luisiturrios/gowin"

	"deng/src/github.com/getlantern/golog"
)

const (
	runDir = `Software\Microsoft\Windows\CurrentVersion\Run`
)

var (
	log = golog.LoggerFor("launcher")
)

func CreateLaunchFile(autoLaunch bool) {
	var startupCommand string

	lanternPath, err := osext.Executable()
	if err != nil {
		log.Errorf("Could not get Lantern directory path: %q", err)
		return
	}

	if autoLaunch {
		// Start Lantern normally.
		startupCommand = fmt.Sprintf(`"%s" -startup`, lanternPath)
	} else {
		// Just clear stored proxy settings and quit.
		startupCommand = fmt.Sprintf(`"%s" -clear-proxy-settings`, lanternPath)
	}

	err = gowin.WriteStringReg("HKCU", runDir, "Lantern", startupCommand)
	if err != nil {
		log.Errorf("Error setting Lantern auto-start registry key: %q", err)
	}
}
