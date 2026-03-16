//go:build !windows

package system

import (
	"os/exec"
)

// OpenBrowser launches the default browser on macOS/Linux
func OpenBrowser(url string) error {
	// On macOS, 'open' is the standard command
	return exec.Command("open", url).Start()
}
