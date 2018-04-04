package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// See https://medium.com/@skdomino/watch-this-file-watching-in-go-5b5a247cf71f for inspiration

func createWatch(path string, call Call, callPipe chan<- Call) (*fsnotify.Watcher, error) {
	// Creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return nil, err
	}

	// Link events to command calls pushed to the callPipe
	go func() {
		for {
			select {
			// watch for events
			case <-watcher.Events:
				callPipe <- call

			// watch for errors
			case watchError := <-watcher.Errors:
				log.Println("ERROR: ", watchError.Error())
			}
		}
	}()

	// Declare the folder to watch, and return the event pipe
	if err := watcher.Add(path); err != nil {
		log.Println("ERROR: ", err.Error())
	}

	return watcher, nil
}

func generateFolderWatchCommands(s *Settings, callPipe chan<- Call) []*fsnotify.Watcher {
	watchers := make([]*fsnotify.Watcher, 0)

	for f := range s.FolderWatch {
		// Create the folder watchs according to settings
		if watcher, err := createWatch(s.FolderWatch[f].FolderToWatch, s.FolderWatch[f].CommandToTrigger, callPipe); err != nil {
			log.Println("ERROR: Failed to create folder watch")
		} else {
			watchers = append(watchers, watcher)
		}
	}
	return watchers
}
