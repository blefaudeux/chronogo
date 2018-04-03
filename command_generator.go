package main

import (
	"log"
	"time"
)

func generateCommands(s *Settings, db *DB, callPipe chan<- Call) {

	for {
		// Go through the commands
		log.Println("** Handling hourly commands ***")

		for c := range s.TimedCommands.Hourly {
			call := s.TimedCommands.Hourly[c]

			// Fetch the last time this command was called:
			if db.startHourly(call.hash()) {
				callPipe <- call
			} else {
				log.Println("Skipping ", call.hash(), ", already called")
			}

		}

		log.Println("** Handling daily commands ***")
		for c := range s.TimedCommands.Daily {
			call := s.TimedCommands.Daily[c]

			if db.startDaily(call.hash()) {
				callPipe <- call
			} else {
				log.Println("Skipping ", call.hash(), ", already called")
			}
		}

		log.Println("** Handling weekly commands ***")
		for c := range s.TimedCommands.Weekly {
			call := s.TimedCommands.Weekly[c]

			if db.startWeekly(call.hash()) {
				callPipe <- call
			} else {
				log.Println("Skipping ", call.hash(), ", already called")
			}
		}

		// Keep going, every hour
		log.Println("-- Sleeping for a while -- ")
		time.Sleep(time.Hour)
	}
}
