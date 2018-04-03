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

	// Start the master/slave command handling
	// -the command generator populates the list of commands to execute
	// -the igniter goes through the commands and starts them
	commandPipe := make(chan Call)
	go generateCommands(&s, &db, commandPipe)
	unstackCommands(&db, commandPipe)

	// (DEBUG) Expose what has been stored in the DB
	log.Println(" --- DEBUG ---")
	keys := db.keys()
	for k := range keys {
		lastCall, _ := db.loadTime(k)
		log.Println("Command: ", k)
		log.Println("*** was last called in ", lastCall.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	}
}
