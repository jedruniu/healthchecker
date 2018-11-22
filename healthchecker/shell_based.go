package healthchecker

import (
	"fmt"
	"os/exec"
	"strings"
)

type shellBasedHealthCheck struct {
	cmd []string
}

func NewShellBased(cmd string) SingleChecker {
	return &shellBasedHealthCheck{strings.Split(cmd, " ")}
}

func (hc *shellBasedHealthCheck) SingleCheck() bool {
	cmd := exec.Command(hc.cmd[0], hc.cmd[1:]...)
	if err := cmd.Start(); err != nil {
		fmt.Printf("could not run command, err: %v\n", err)
		return false
	}
	return cmd.Wait() == nil
}
