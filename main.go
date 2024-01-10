package main

import (
	"fmt"
	macos "main/macOS"
	"main/webbrowser"
)

func main() {

	// Open a web browser and search for "golang"
	err := webbrowser.Search("Hello World")
	if err != nil {
		fmt.Println("Error opening web browser:", err)
	}

	// Send a macOS notification with the title "Hello" and the message "World"
	err = macos.CreateNotification("Hello", "World")
	if err != nil {
		fmt.Println("Error creating notification:", err)
	}
}
