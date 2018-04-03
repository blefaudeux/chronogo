package main

import "log"

func unstackCommands(db *DB, callPipe <-chan Call) {
	for {
		call, more := <-callPipe
		if more {
			if cmd, err := startProcess(call.Command, call.Args); err == nil {
				// Start a go routine, save asynchronously when the command is done
				go func() {
					if err := cmd.Wait(); err == nil {
						db.storeTime(call.hash())
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
