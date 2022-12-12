package beeep

import (
	"testing"
)

func TestNotify(t *testing.T) {
	err := Notify("Notify title", "Message body", "beeep", "assets/information.png")
	if err != nil {
		t.Error(err)
	}
}
