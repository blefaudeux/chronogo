package main

import (
	"strings"
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

	Log.Println("Opening the DB in:", path)

	return DB{diskv.New(diskv.Options{
		BasePath:     path,
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})}
}

func (d *DB) store(key string, val []byte) {
	d.disk.Write(handleKey(key), val)
}

func (d *DB) storeTime(key string) {
	// Store the current time alongside the key
	t, _ := time.Now().MarshalJSON()
	if err := d.disk.Write(handleKey(key), t); err != nil {
		Log.Println("ERROR: DB (Write):", handleKey(key), err.Error())
	}
}

func (d *DB) load(key string) []byte {
	// Read the value back out of the store.
	var value []byte
	var err error

	if value, err = d.disk.Read(handleKey(key)); err != nil {
		Log.Println("ERROR: DB (Read):", handleKey(key), err.Error())
	}
	return value
}

func handleKey(s string) string {
	return strings.Replace(strings.Replace(s, "\\", "", -1), "/", "", -1)
}

func (d *DB) loadTime(key string) (time.Time, error) {
	// If the key is not present, the default time is returned along with an error
	value, err := d.disk.Read(handleKey(key))
	var t time.Time
	if err != nil {
		return t, err
	}

	t.UnmarshalJSON(value)
	return t, nil
}

func (d *DB) erase(key string) {
	// Erase the key+value from the DB.
	if err := d.disk.Erase(handleKey(key)); err != nil {
		Log.Println("ERROR: DB (Erase):", handleKey(key), err.Error())
	}
}

func (d *DB) keys() <-chan string {
	return d.disk.Keys(nil)
}

// Easier getters to check whether those commands should be started or not
func (d *DB) startHourly(key string) bool {
	// Try to load this key, if not present then it can be started
	lastCall, err := d.loadTime(key)
	if err != nil {
		Log.Println("Command: ", key, " was never called")
		return true
	}

	return int(time.Since(lastCall).Hours()) > 0
}

func (d *DB) startDaily(key string) bool {
	// Try to load this key, if not present then it can be started
	lastCall, err := d.loadTime(key)
	if err != nil {
		Log.Println("Command: ", key, " was never called")
		return true
	}

	return time.Since(lastCall).Hours() > 24.
}

func (d *DB) startWeekly(key string) bool {
	// Try to load this key, if not present then it can be started
	lastCall, err := d.loadTime(key)
	if err != nil {
		Log.Println("Command: ", key, " was never called")
		return true
	}

	return time.Since(lastCall).Hours() > 24*7
}

func (d *DB) startMonthly(key string) bool {
	// Try to load this key, if not present then it can be started
	lastCall, err := d.loadTime(key)
	if err != nil {
		Log.Println("Command: ", key, " was never called")
		return true
	}

	return time.Since(lastCall).Hours() > 24*7*30
}
