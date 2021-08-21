// +build !darwin

package pac

import (
	"deng/src/github.com/getlantern/byteexec"
)

func ensureElevatedOnDarwin(be *byteexec.Exec, prompt string, iconFullPath string) (err error) {
	return nil
}
