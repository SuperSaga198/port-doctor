//go:build darwin

package osutil

import (
	"os/exec"
	"strconv"
	"strings"
)

func FindProcessByPort(port int) (int, string, error) {
	// Use lsof to find the PID listening on the port
	cmd := exec.Command("lsof", "-t", "-i:"+strconv.Itoa(port))
	out, err := cmd.Output()
	if err != nil {
		return 0, "", nil // No process found
	}

	pidStr := strings.TrimSpace(string(out))
	// lsof might return multiple lines if multiple states, take the first PID
	lines := strings.Split(pidStr, "\n")
	if len(lines) == 0 || lines[0] == "" {
		return 0, "", nil
	}

	pid, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, "", err
	}

	// Get process command name
	nameCmd := exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "comm=")
	nameOut, err := nameCmd.Output()
	name := "unknown"
	if err == nil {
		name = strings.TrimSpace(string(nameOut))
	}

	return pid, name, nil
}

func KillProcess(pid int) error {
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	return cmd.Run()
}