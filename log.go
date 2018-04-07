package main

import (
	"log"
	"os"
)

// See github/border for inspiration

// NewLog initializes the log file pipe
func NewLog(logpath string) *log.Logger {
	println("LogFile: " + logpath)
	file, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}
	return log.New(file, "", log.LstdFlags|log.Lshortfile)
}
