package utils

// Defines the utility package that provides helper functions.

import (
	"os/exec" // Used to execute external commands.
	"runtime" // To determine the operating system runtime.
)

// OpenBrowser tries to open the specified URL in the default web browser.
func OpenBrowser(url string) error {
	var cmd string    // Holds the command to open the browser.
	var args []string // Holds arguments for the command.

	switch runtime.GOOS { // Checks the operating system.
	case "windows":
		cmd = "cmd"                    // Command prompt for Windows.
		args = []string{"/c", "start"} // 'start' opens the default browser.
	case "darwin":
		cmd = "open" // 'open' is used on macOS to open the browser.
	default:
		cmd = "xdg-open" // Linux's systems use 'xdg-open' to open the default browser.
	}
	args = append(args, url)                  // Append the URL to the command arguments.
	return exec.Command(cmd, args...).Start() // Execute the command to open the browser.
}
