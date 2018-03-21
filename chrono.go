package main

import "fmt"
import "flag"

func main() {
	settingsPath := flag.String("settingsPath", "~/.chronogo", "Path to the settings file")
	waitForConnection := flag.Bool("waitForConnection", true, "Wait for internet connection before starting the jobs")
	flag.Parse()

	fmt.Printf("Settings path: %s\nWait for connection: %t", *settingsPath, *waitForConnection)
}
