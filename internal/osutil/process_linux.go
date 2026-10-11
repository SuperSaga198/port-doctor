//go:build linux

package osutil

import (
	"os/exec"
	"strconv"
	"strings"
)

func FindProcessByPort(port int) (int, string, error) {
	// Use ss or netstat to find PID
	cmd := exec.Command("ss", "-lptn", "sport = :"+strconv.Itoa(port))
	out, err := cmd.Output()
	if err != nil {
		// Fallback to lsof if ss fails
		cmd = exec.Command("lsof", "-t", "-i:"+strconv.Itoa(port))
		out, err = cmd.Output()
		if err != nil {
			return 0, "", nil
		}
	}

	// Simple parser for lsof / ss output lookup
	pidStr := strings.TrimSpace(string(out))
	lines := strings.Split(pidStr, "\n")
	if len(lines) == 0 || lines[0] == "" {
		return 0, "", nil
	}

	// If ss was used, extract PID, or if lsof was used, it's already PID
	pid, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, "", nil
	}

	// Get command name
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