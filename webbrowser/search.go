// File: webbrowser/search.go

package webbrowser

import (
	"fmt"
	"net/url"
	"os/exec"
)

// Search opens the default web browser and performs a search for the given query
func Search(query string) error {
	escapedQuery := url.QueryEscape(query)
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", escapedQuery)
	cmd := exec.Command("open", searchURL) // 'open' for macOS, use 'xdg-open' for Linux, 'start' for Windows
	return cmd.Run()
}
