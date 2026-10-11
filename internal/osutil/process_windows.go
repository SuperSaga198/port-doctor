//go:build windows

package osutil

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func FindProcessByPort(port int) (int, string, error) {
	// Use netstat to find PID listening on port
	cmd := exec.Command("netstat", "-ano")
	out, err := cmd.Output()
	if err != nil {
		return 0, "", err
	}

	lines := strings.Split(string(out), "\n")
	target := fmt.Sprintf(":%d", port)
	var pid int

	for _, line := range lines {
		if strings.Contains(line, target) && strings.Contains(line, "LISTENING") {
			fields := strings.Fields(line)
			if len(fields) > 0 {
				pidStr := fields[len(fields)-1]
				p, err := strconv.Atoi(pidStr)
				if err == nil {
					pid = p
					break
				}
			}
		}
	}

	if pid == 0 {
		return 0, "", nil
	}

	// Get process executable name via tasklist
	tCmd := exec.Command("tasklist", "/fi", fmt.Sprintf("PID eq %d", pid), "/fo", "CSV", "/nh")
	tOut, err := tCmd.Output()
	name := "unknown"
	if err == nil {
		parts := strings.Split(string(tOut), ",")
		if len(parts) > 0 {
			name = strings.Trim(parts[0], "\"")
		}
	}

	return pid, name, nil
}

func KillProcess(pid int) error {
	cmd := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
	return cmd.Run()
}