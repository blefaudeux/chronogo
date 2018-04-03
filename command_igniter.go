package main

import (
	"log"
)

func unstackCommands(db *DB, callPipe <-chan Call) {

	for {
		call, stillGood := <-callPipe
		if stillGood {
			if cmd, err := startProcess(call.Command, call.Args); err == nil {
				// Start a go routine, save asynchronously when the command is done
				go func() {
					if err := cmd.Wait(); err == nil {
						db.storeTime(call.hash())
						log.Println("Command", call.hash(), "completed")
					} else {
						log.Println("Command", call.hash(), "failed", "Error: ", err.Error())
					}
				}()
			}
		} else {
			break
		}

	}
}
