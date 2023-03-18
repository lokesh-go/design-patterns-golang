package singleton

import (
	"log"
	"sync"
)

type single struct{}

// Global instance
var singleInstance *single
var lock = &sync.Mutex{}

// GetInstance ...
func GetInstance() *single {
	if singleInstance == nil {
		// Lock
		lock.Lock()
		defer lock.Unlock()

		// Creating thread safe singleton object
		// Handling for multiple thread comes at the same time
		// Ensure that the instance hasn't yet been initialised by another thread
		// while one has been waiting for lock's release
		if singleInstance == nil {
			log.Println("Creating single instance now")
			singleInstance = &single{}
		} else {
			log.Println("Single instance already created")
		}
	} else {
		log.Println("Single instance already created")
	}
	return singleInstance
}
