package main

import (
	"flag"
	"log"
)

func main() {
	settingsPath := flag.String("settingsPath", "~/.chronogo", "Path to the settings file")
	flag.Parse()

	log.Println("Settings path:", *settingsPath)

	// s := loadSettings(*settingsPath)
	var s Settings
	s.dumpToFile(*settingsPath)
}
