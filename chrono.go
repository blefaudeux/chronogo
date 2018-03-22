package main

import (
	"flag"
	"log"
)

func main() {
	settingsPath := flag.String("settingsPath", "~/.chronogo", "Path to the settings file")
	initSettings := flag.Bool("initSettings", false, "This will reset the settings in 'settingsPath'")
	flag.Parse()

	if *initSettings {
		log.Println("Reset the settings file in ", *settingsPath)
		var s Settings
		s.dumpToFile(*settingsPath)
		return
	}

	log.Println("Loading the settings in ", *settingsPath)
	s := loadSettings(*settingsPath)

	log.Println("** Starting daily commands")
	for c := range s.TimedCommands.Daily {
		// TODO: Make that non-blocking
		startProcess(s.TimedCommands.Daily[c].Command, s.TimedCommands.Daily[c].Args)
	}
}
