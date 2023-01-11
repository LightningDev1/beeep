//go:build js
// +build js

package beeep

// Alert displays a desktop notification and plays a beep.
func Alert(title, message, appID, appIcon string) error {
	if err := Notify(title, message, appID, appIcon); err != nil {
		return err
	}
	return Beep(DefaultFreq, DefaultDuration)
}
