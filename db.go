package main

import (
	"log"
	"time"

	"github.com/peterbourgon/diskv"
)

// See https://github.com/peterbourgon/diskv
// Thin wrapper around diskv

// DB describes the on-disk persistent data base
type DB struct {
	disk *diskv.Diskv
}

func initNew(path string) DB {
	// Simplest transform function: put all the data files into the base dir.
	flatTransform := func(s string) []string { return []string{} }

	// DEMO: Initialize a new diskv store, with a 1MB cache.
	return DB{diskv.New(diskv.Options{
		BasePath:     path,
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})}
}

func (d *DB) store(key string, val []byte) {
	d.disk.Write(key, val)
}

func (d *DB) storeTime(key string) {
	// Store the current time alongside the key
	t, _ := time.Now().MarshalJSON()
	d.disk.Write(key, t)
}

func (d *DB) load(key string) []byte {
	// Read the value back out of the store.
	value, _ := d.disk.Read(key)
	return value
}

func (d *DB) loadTime(key string) (time.Time, error) {
	// If the key is not present, the default time is returned along with an error
	value, err := d.disk.Read(key)
	var t time.Time
	if err != nil {
		return t, err
	}

	t.UnmarshalJSON(value)
	return t, nil
}

func (d *DB) erase(key string) {
	// Erase the key+value from the DB.
	d.disk.Erase(key)
}

func (d *DB) keys() <-chan string {
	return d.disk.Keys(nil)
}

// Easier getters to check whether those commands should be started or not
func (d *DB) startHourly(key string) bool {
	// Try to load this key, if not present then it can be started
	lastCall, err := d.loadTime(key)
	if err != nil {
		log.Println("Command: ", key)
		log.Println("*** was never called")
		return true
	}

	log.Println("Command: ", key)
	log.Println("*** was last called in ", lastCall.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))

	return time.Since(lastCall).Hours() > 0
}

func (d *DB) startDaily(key string) bool {
	// Try to load this key, if not present then it can be started
	lastCall, err := d.loadTime(key)
	if err != nil {
		log.Println("Command: ", key)
		log.Println("*** was never called")
		return true
	}

	log.Println("Command: ", key)
	log.Println("*** was last called in ", lastCall.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))

	return time.Since(lastCall).Hours() > 24.
}
