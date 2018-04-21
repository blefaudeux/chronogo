package main

func unstackCommands(dbDone *DB, dbStarted *DB, callPipe <-chan Call) {

	for {
		call, stillGood := <-callPipe
		if stillGood {
			if cmd, err := startProcess(call.Command, call.Args); err == nil {
				// Start a go routine, save asynchronously when the command is done
				go func() {
					Log.Println("Starting", call.hash())
					dbStarted.storeTime(call.hash())

					if err := cmd.Wait(); err == nil {
						dbDone.storeTime(call.hash())
						Log.Println("Command", call.hash(), "completed")
					} else {
						Log.Println("Command", call.hash(), "failed", "Error: ", err.Error())
					}
				}()
			}
		} else {
			break
		}

	}
}
