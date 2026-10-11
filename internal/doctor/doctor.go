package doctor

import (
	"fmt"
	"github.com/SuperSaga198/port-doctor/internal/osutil"
)

func CheckPort(port int) (int, string, error) {
	return osutil.FindProcessByPort(port)
}

func FreePort(port int, force bool) error {
	pid, name, err := CheckPort(port)
	if err != nil {
		return err
	}
	if pid == 0 {
		fmt.Printf("Port %d is already free!\n", port)
		return nil
	}

	fmt.Printf("Found process '%s' (PID: %d) occupying port %d.\n", name, pid, port)

	if !force {
		fmt.Print("Do you want to terminate this process? (y/N): ")
		var answer string
		fmt.Scanln(&answer)
		if answer != "y" && answer != "Y" {
			fmt.Println("Operation cancelled.")
			return nil
		}
	}

	err = osutil.KillProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to kill process PID %d: %v", pid, err)
	}

	fmt.Printf("SUCCESS: Process %d terminated. Port %d is now free!\n", pid, port)
	return nil
}