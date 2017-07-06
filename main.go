package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

var (
	service = flag.String("s", "", "A comma-separated list of services to monitor")
)

func serviceCheck(s string) string {
	//Command to check if systemd service is active
	cmdName := "/bin/systemctl"
	cmdArgs := []string{"is-active", s}

	cmdOut, _ := exec.Command(cmdName, cmdArgs...).Output()

	isActive := strings.TrimSpace(string(cmdOut))
	if isActive == "active" {
		fmt.Printf("%s is active\n", s)
	} else {
		fmt.Printf("%s is not active\n", s)
	}
	return isActive
}

func main() {
	flag.Parse()
	//split up based on comma
	serviceSlice := strings.Split(*service, ",")

	if *service == "" {
		fmt.Println("You need to define a comma-separated list of services to monitor.")
	} else {
		for i := range serviceSlice {
			serviceCheck(serviceSlice[i])
		}
	}
}
