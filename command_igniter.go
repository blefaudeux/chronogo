package main

import "log"

func unstackCommands(db *DB, callPipe <-chan Call) {
	for {
		call, more := <-callPipe
		if more {
			log.Println("Starting command", call.Command)
			if err := startProcess(call.Command, call.Args); err == nil {
				db.storeTime(call.hash())
			}
		} else {
			break
		}

	}
}
