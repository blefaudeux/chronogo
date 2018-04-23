package main

import (
	"time"
)

func generateTimedCommands(s *Settings, dbDone *DB, dbStarted *DB, callPipe chan<- Call) {

	for {
		// Go through the commands
		Log.Println("** Handling hourly commands ***")
		for c := range s.TimedCommands.Hourly {
			call := s.TimedCommands.Hourly[c]

			// Fetch the last time this command completed:
			if dbDone.startHourly(call.hash()) && dbStarted.startHourly(call.hash()) {
				callPipe <- call
			} else {
				Log.Println("Skipping ", call.hash(), ", already called")
			}
		}
		Log.Println("-> Hourly commands handled")

		Log.Println("** Handling daily commands ***")
		for c := range s.TimedCommands.Daily {
			call := s.TimedCommands.Daily[c]

			if dbDone.startDaily(call.hash()) && dbStarted.startDaily(call.hash()) {
				callPipe <- call
			} else {
				Log.Println("Skipping ", call.hash(), ", already called")
			}
		}
		Log.Println("-> Daily commands handled")

		Log.Println("** Handling weekly commands ***")
		for c := range s.TimedCommands.Weekly {
			call := s.TimedCommands.Weekly[c]

			if dbDone.startWeekly(call.hash()) && dbStarted.startWeekly(call.hash()) {
				callPipe <- call
			} else {
				Log.Println("Skipping ", call.hash(), ", already called")
			}
		}
		Log.Println("-> Weekly commands handled")

		Log.Println("** Handling monthly commands ***")
		for c := range s.TimedCommands.Monthly {
			call := s.TimedCommands.Monthly[c]

			if dbDone.startMonthly(call.hash()) && dbStarted.startMonthly(call.hash()) {
				callPipe <- call
			} else {
				Log.Println("Skipping ", call.hash(), ", already called")
			}
		}
		Log.Println("-> Monthly commands handled")

		// Keep going, every hour
		Log.Println("-- Chronogo is sleeping for a while --")
		time.Sleep(time.Hour)
	}
}
