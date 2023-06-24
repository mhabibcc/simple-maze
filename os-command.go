package main

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() {
	cmd := exec.Command(clearCommand())
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearCommand() string {
	switch runtime.GOOS {
	case "windows":
		return "cls"
	case "linux", "darwin":
		return "clear"
	default:
		panic("unsupported platform")
	}
}
