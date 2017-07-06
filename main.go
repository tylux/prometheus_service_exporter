package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var (
	service = flag.String("s", "", "The name of the systemd service you want to monitor")
)

func serviceCheck(s string) string {
	//Command to check if systemd service is active
	cmdName := "systemctl is-active"
	cmdArgs := []string{s}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}

	isActive := string(cmdOut)

	if isActive == "active" {
		fmt.Printf("%s is active", s)
	} else {
		fmt.Printf("%s is not active", s)
	}
	return isActive
}

func main() {
	flag.Parse()

	if *service == "" {
		fmt.Println("You need to define a service to monitor..")
	} else {
		serviceCheck(*service)
	}

}
