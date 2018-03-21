package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// Settings : holds all the commands and triggers used by chronogo
type Settings struct {
	waitForInternetConnection bool
	timedCommands             struct {
		hourly  []string
		daily   []string
		weekly  []string
		monthly []string
	}
	folderWatchCommands []struct {
		folderToWatch    string
		commandToTrigger string
	}
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

func loadSettings(path string) Settings {
	log.Println("Loading settings from:", path)

	raw, err := ioutil.ReadFile(path)
	check(err)

	var s Settings
	json.Unmarshal(raw, &s)
	return s
}

func (s *Settings) dumpToFile(path string) {
	log.Println("Saving settings in:", path)

	// Sanitize path
	if path[:2] == "~/" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		path = filepath.Join(dir, path[2:])
	}

	f, err := os.Create(path)
	check(err)
	defer f.Close()

	_, err = f.Write([]byte(toJSON(s)))
	check(err)
}
