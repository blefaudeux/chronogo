package main

import "os/exec"

func startProcess(command string, args []string) (*exec.Cmd, error) {
	cmd := exec.Cmd{Path: command, Args: args}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// FIXME: will cmd go out of scope ?
	return &cmd, nil
}
