package main

import (
	"fmt"

	app "gioui.org/app"

	"gioui.org/font/gofont"
)

func main() {
	gofont.Register()
	win, err := createWindow(transactionsPage)
	if err != nil {
		fmt.Printf("Could not initialize window: %s\ns", err)
		return
	}
	go func(win *window) {
		win.loop()
	}(win)

	app.Main()
}
