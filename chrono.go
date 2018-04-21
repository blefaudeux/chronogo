package main

import (
	"flag"
	"log"
)

var (
	// Log is the file pipe which will receive all the logs
	Log *log.Logger
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

	// Initialize the logger
	Log = NewLog(s.LogPath)
	Log.Println("Settings loaded from", *settingsPath)

	// Load the DB
	dbDone := initNew(s.DBPath)
	dbStarted := initNew(s.DBPath + "_start")

	// Start the master/slave command handling
	// -the command generators populates the list of commands to execute
	// -the igniter goes through the commands and starts them
	commandPipe := make(chan Call)

	// - the recurrent commands
	go generateTimedCommands(&s, &dbDone, &dbStarted, commandPipe)

	// - the commands based on folder triggers
	watchers := generateFolderWatchCommands(&s, commandPipe)
	for w := range watchers {
		defer watchers[w].Close()
	}
	Log.Println("Folder watch initialized, ", len(watchers), " of them in flight")

	// - this one will block and execute the incoming commands
	unstackCommands(&dbDone, &dbStarted, commandPipe)
}
