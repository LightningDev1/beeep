package beeep

func ExampleBeep() {
	Beep(DefaultFreq, DefaultDuration)
}

func ExampleNotify() {
	Notify("Title", "MessageBody", "beeep", "assets/information.png")
}

func ExampleAlert() {
	Alert("Title", "MessageBody", "beeep", "assets/warning.png")
}
