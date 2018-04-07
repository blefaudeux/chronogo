package main

import (
	"time"
)

func generateTimedCommands(s *Settings, db *DB, callPipe chan<- Call) {

	for {
		// Go through the commands
		Log.Println("** Handling hourly commands ***")

		for c := range s.TimedCommands.Hourly {
			call := s.TimedCommands.Hourly[c]

			// Fetch the last time this command was called:
			if db.startHourly(call.hash()) {
				callPipe <- call
			} else {
				Log.Println("Skipping ", call.hash(), ", already called")
			}
		}
		Log.Println("-> Hourly commands handled")

		Log.Println("** Handling daily commands ***")
		for c := range s.TimedCommands.Daily {
			call := s.TimedCommands.Daily[c]

			if db.startDaily(call.hash()) {
				callPipe <- call
			} else {
				Log.Println("Skipping ", call.hash(), ", already called")
			}
		}
		Log.Println("-> Daily commands handled")

		Log.Println("** Handling weekly commands ***")
		for c := range s.TimedCommands.Weekly {
			call := s.TimedCommands.Weekly[c]

			if db.startWeekly(call.hash()) {
				callPipe <- call
			} else {
				Log.Println("Skipping ", call.hash(), ", already called")
			}
		}
		Log.Println("-> Weekly commands handled")

		// Keep going, every hour
		Log.Println("-- Sleeping for a while -- ")
		time.Sleep(time.Hour)
	}
}
