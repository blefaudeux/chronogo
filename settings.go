package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// Call describes everything needed to execute a given command
type Call struct {
	Command string
	Args    []string
}

func (c *Call) hash() string {
	result := c.Command
	for a := range c.Args {
		result += "_" + c.Args[a]
	}
	return result
}

// FolderTriggeredCall describes a command which is executed when a given folder is touched
type FolderTriggeredCall struct {
	FolderToWatch    string
	CommandToTrigger Call
}

// TimedCalls describes a collection of commands which are executed on a timely basis
type TimedCalls struct {
	Hourly  []Call
	Daily   []Call
	Weekly  []Call
	Monthly []Call
}

// Settings : holds all the commands and triggers used by chronogo
type Settings struct {
	TimedCommands       TimedCalls
	FolderWatch         []FolderTriggeredCall
	DBPath              string
	MaxCommandsInFlight int
}

func defaultSettings() Settings {
	hourly := []Call{Call{"echo", []string{"this command", "will be run", "every hour"}}}
	daily := []Call{Call{"echo", []string{"this command", "will be run", "every day"}}}
	weekly := []Call{Call{"echo", []string{"this command", "will be run", "every week"}}}
	monthly := []Call{Call{"echo", []string{"this command", "will be run", "every month"}}}

	defaultTimed := TimedCalls{hourly, daily, weekly, monthly}
	defaultFolder := []FolderTriggeredCall{FolderTriggeredCall{"/home/username", Call{"echo", []string{"this command is triggered when the folder is changed"}}}}

	return Settings{defaultTimed, defaultFolder, "chronoDB", 2}
}

func (s *Settings) toString() string {
	// Load the parsed settings here
	return toJSON(s)
}

func toJSON(s *Settings) string {
	bytes, err := json.Marshal(*s)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

// IO helpers

func check(e error) {
	if e != nil {
		log.Println(e.Error())
		panic(e)
	}
}

func sanitizePath(path string) string {
	if path[:2] == "~/" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		path = filepath.Join(dir, path[2:])
	}
	return path
}

func loadSettings(path string) Settings {
	raw, err := ioutil.ReadFile(sanitizePath(path))
	check(err)

	var s Settings
	json.Unmarshal(raw, &s)

	// Handle a reasonable default for the max concurrent commands
	if s.MaxCommandsInFlight == 0 {
		s.MaxCommandsInFlight = 3
	}
	return s
}

func (s *Settings) dumpToFile(path string) {
	// Sanitize path
	f, err := os.Create(sanitizePath(path))
	check(err)
	defer f.Close()

	_, err = f.Write([]byte(toJSON(s)))
	check(err)
}
