//go:build windows

package system

import (
	"fmt"
	"os/exec"
)

// OpenBrowser launches the default browser to the given URL using Windows rundll32
func OpenBrowser(url string) error {
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to open browser: %w", err)
	}
	return nil
}
