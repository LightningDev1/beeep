//go:build windows && !linux && !freebsd && !netbsd && !openbsd && !darwin && !js
// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package beeep

import (
	toast "github.com/go-toast/toast"
)

// Alert displays a desktop notification and plays a default system sound.
func Alert(title, message, appID, appIcon string) error {
	if isWindows10 {
		note := toastNotification(title, message, appID, pathAbs(appIcon))
		note.Audio = toast.Default
		return note.Push()
	}

	if err := Notify(title, message, appID, appIcon); err != nil {
		return err
	}
	return Beep(DefaultFreq, DefaultDuration)
}
