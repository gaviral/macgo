// File: macOS/notification.go

package macos

import (
	"fmt"
	"os/exec"
)

// CreateNotification displays a macOS notification with the given title and message
func CreateNotification(title, message string) error {
	appleScript := fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)
	cmd := exec.Command("osascript", "-e", appleScript)
	return cmd.Run()
}
