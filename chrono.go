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
		s := defaultSettings()
		s.dumpToFile(*settingsPath)
		return
	}

	log.Println("Loading the settings in ", *settingsPath)
	s := loadSettings(*settingsPath)

	// Load the DB
	db := initNew(s.DBPath)

	// Start the master/slave command handling
	// -the command generators populates the list of commands to execute
	// -the igniter goes through the commands and starts them
	commandPipe := make(chan Call)

	// - the recurrent commands
	go generateTimedCommands(&s, &db, commandPipe)

	// - the commands based on folder triggers
	watchers := generateFolderWatchCommands(&s, commandPipe)
	for w := range watchers {
		defer watchers[w].Close()
	}
	log.Println("Folder watch initialized, ", len(watchers), " of them in flight")

	// - this one will block and execute the incoming commands
	unstackCommands(&db, commandPipe)
}
