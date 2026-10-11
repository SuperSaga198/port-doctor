package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/SuperSaga198/port-doctor/internal/doctor"
)

func Execute() {
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	freeCmd := flag.NewFlagSet("free", flag.ExitOnError)
	forceFlag := freeCmd.Bool("f", false, "Force kill without confirmation prompt")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "check":
		checkCmd.Parse(os.Args[2:])
		args := checkCmd.Args()
		if len(args) == 0 {
			fmt.Println("Error: Please specify a port number.")
			os.Exit(1)
		}
		port, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Invalid port number.")
			os.Exit(1)
		}

		pid, name, err := doctor.CheckPort(port)
		if err != nil {
			fmt.Printf("Error checking port: %v\n", err)
			os.Exit(1)
		}
		if pid == 0 {
			fmt.Printf("Port %d is CLEAR (Available).\n", port)
		} else {
			fmt.Printf("Port %d is OCCUPIED by process '%s' (PID: %d).\n", port, name, pid)
		}

	case "free":
		freeCmd.Parse(os.Args[2:])
		args := freeCmd.Args()
		if len(args) == 0 {
			fmt.Println("Error: Please specify a port number.")
			os.Exit(1)
		}
		port, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Invalid port number.")
			os.Exit(1)
		}

		err = doctor.FreePort(port, *forceFlag)
		if err != nil {
			fmt.Printf("Error freeing port: %v\n", err)
			os.Exit(1)
		}

	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("port-doctor - Network & Process Diagnostics CLI")
	fmt.Println("\nUsage:")
	fmt.Println("  port-doctor check <port>      Diagnose what process is using a port")
	fmt.Println("  port-doctor free <port> [-f]  Inspect and terminate the process holding a port")
}