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

// Settings : holds all the commands and triggers used by chronogo
type Settings struct {
	WaitForInternetConnection bool
	TimedCommands             struct {
		Hourly  []Call
		Daily   []Call
		Weekly  []Call
		Monthly []Call
	}
	FolderWatchCommands []struct {
		FolderToWatch    string
		CommandToTrigger Call
	}
	DBPath              string
	MaxCommandsInFlight int
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
