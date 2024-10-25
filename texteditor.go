package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func openEditor(fileName string) error {
	var cmd *exec.Cmd

	// Determine the operating system
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("notepad", fileName)
	case "darwin": // macOS
		cmd = exec.Command("open", fileName)
	case "linux":
		cmd = exec.Command("nano", fileName) // Using nano instead of xdg-open
	default:
		return fmt.Errorf("unsupported operating system")
	}

	// Start the editor
	return cmd.Start()
}
