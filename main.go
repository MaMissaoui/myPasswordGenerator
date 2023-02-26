package main

import "my-password-generator/gui"

func main() {
	mainScreen := gui.MakeMainScreen()
	gui.Start(mainScreen)
}
