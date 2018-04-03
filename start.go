package main

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

func stdOutToLog(pipe io.ReadCloser) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}

func startProcess(command string, args []string) (*exec.Cmd, error) {
	log.Println("Starting command *", command, "* with the arguments ", args)

	cmd := exec.Command(command, args...)

	// Pipe the stdout to the log
	if stdOut, err := cmd.StdoutPipe(); err == nil {
		go stdOutToLog(stdOut)
	}

	// Start the command, don't wait for exit (for now ?)
	if err := cmd.Start(); err != nil {
		log.Println(err.Error())
		return cmd, err
	}

	return cmd, nil
}
