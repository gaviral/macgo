package main

import (
	"fmt"
	macos "main/macOS"
	"main/openaiapi"
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

	// Create an image using the OpenAI API
	prompt := "A cute baby sea otter"
	n := 1
	size := "1024x1024"

	response, err := openaiapi.CreateImage(prompt, n, size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response from OpenAI:")
	for _, img := range response.Data {
		fmt.Println("Image URL:", img.URL)
	}

}
