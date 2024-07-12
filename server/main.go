package main

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func main() {
	// Command to run in the pseudo terminal
	c := exec.Command("bash")

	// Start the pseudo terminal
	ptmx, err := pty.Start(c)
	if err != nil {
		log.Fatalf("pty.Start: %v", err)
	}
	defer ptmx.Close()

	// Transfer input/output between the terminal and the pty
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)
}
