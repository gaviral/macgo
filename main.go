package main

import (
	"fmt"
	macos "main/macOS"
)

func main() {
	err := macos.CreateNotification("OpenAI Response", "test")
	if err != nil {
		fmt.Println("Error creating notification:", err)
	}
}
