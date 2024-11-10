package shell

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

type ShellType string
const (
	ShellTypeBash ShellType = "bash"
	ShellTypeCsh  ShellType = "csh"
	ShellTypeDash ShellType = "dash"
	ShellTypeKsh  ShellType = "ksh"
	ShellTypeSh   ShellType = "sh"
	ShellTypeTcsh ShellType = "tcsh"
	ShellTypeZsh  ShellType = "zsh"
)

func (s ShellType) string() string {
	return string(s)
}

func (s ShellType) isValid() bool {
	return s == ShellTypeBash || s == ShellTypeCsh || s == ShellTypeDash || s == ShellTypeKsh || s == ShellTypeSh || s == ShellTypeTcsh || s == ShellTypeZsh
}

// ExecuteCommand executes a command and returns the output.
func ExecuteCommand(shellType ShellType, command string) (string, error) {
	if !shellType.isValid() {
		return "", fmt.Errorf("invalid shell type: %s", shellType)
	}

	var out bytes.Buffer
	cmdExec := exec.Command(shellType.string(), "-c", command)
	cmdExec.Stdout = &out
	cmdExec.Stderr = &out

	// Start the command without waiting for it to complete
	err := cmdExec.Run()
	if err != nil {
		fmt.Printf("Error starting command '%s': %v\n", command, err)
		return "", err
	}

	return out.String(), nil
}

// ExecuteCommandWait executes a command and waits for it to complete.
func ExecuteCommandWait(shellType ShellType, command string, timeout time.Duration) (string, error) {
	if !shellType.isValid() {
		return "", fmt.Errorf("invalid shell type: %s", shellType)
	}

	var out bytes.Buffer
	cmdExec := exec.Command(shellType.string(), "-i", "-c", command)
	cmdExec.Stdout = &out
	cmdExec.Stderr = &out

	// Start the command and handle errors
	if err := cmdExec.Start(); err != nil {
		fmt.Printf("Error starting command '%s': %v\n", command, err)
		return "", err
	}

	// Create a channel to signal when the command finishes
	done := make(chan error)
	go func() {
		done <- cmdExec.Wait()
	}()

	// Use select to handle timeout or successful command completion
	select {
	case <-time.After(timeout):
		// Kill the process if it times out
		if err := cmdExec.Process.Kill(); err != nil {
			fmt.Printf("Failed to kill command '%s': %v\n", command, err)
			return "", err
		}
		return "", fmt.Errorf("command '%s' timed out", command)

	case err := <-done:
		// If the command completed with an error, return it
		if err != nil {
			fmt.Printf("Command '%s' ended with error: %v\n", command, err)
			return "", err
		}
	}

	return out.String(), nil
}