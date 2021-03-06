// This simple test program prints out whether or not the default browser is
// Edge, which is useful for testing/debugging edgedetect.
package main

import (
	"deng/src/github.com/getlantern/edgedetect"
	"deng/src/github.com/getlantern/golog"
)

var (
	log = golog.LoggerFor("edgedetect.cmd")
)

func main() {
	log.Debugf("Is Edge: %v", edgedetect.DefaultBrowserIsEdge())
}
