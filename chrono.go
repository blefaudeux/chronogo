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

	// Load the DB
	db := initNew(s.DBPath)

	// Go through the commands
	log.Println("** Starting hourly commands")

	for c := range s.TimedCommands.Hourly {
		call := s.TimedCommands.Hourly[c]

		// Fetch the last time this command was called:
		if db.startHourly(call.hash()) {
			if err := startProcess(call.Command, call.Args); err == nil {
				db.storeTime(call.hash())
			}
		} else {
			log.Println("Skipping ", call.hash(), ", already called")
		}

	}

	log.Println("** Starting daily commands")
	for c := range s.TimedCommands.Daily {
		call := s.TimedCommands.Daily[c]

		// Fetch the last time this command was called:
		if db.startDaily(call.hash()) {
			if err := startProcess(call.Command, call.Args); err == nil {
				db.storeTime(call.hash())
			}
		} else {
			log.Println("Skipping ", call.hash(), ", already called")
		}
	}

	// Expose what has been stored in the DB
	log.Println(" --- Foolproofing ---")
	keys := db.keys()
	for k := range keys {
		lastCall, _ := db.loadTime(k)
		log.Println("Command: ", k)
		log.Println("*** was last called in ", lastCall.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	}
}
