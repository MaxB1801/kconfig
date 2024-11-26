package main

import (
	"fmt"
	"os"
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
		cmd := exec.Command("bash", "-c", fmt.Sprintf("nano %s", fileName))
		// Attach the terminal's input/output to the command
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the editor and wait for it to complete
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening nano: %v\n", err)
			os.Exit(1)
		}
		return nil
	default:
		return fmt.Errorf("unsupported operating system")
	}

	// Start the editor
	return cmd.Start()
}
