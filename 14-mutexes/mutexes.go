package main

import "sync"

/*
	Mutexes
		Allow us to lock access to data, ensuring that we can control which goroutines can access
		certain data at which time.
		Go provides built-in mutex feature with sync.Mutex and its 2 methods: .Lock() and .Unlock()

		It's a good practice to put the protected code inside a function, so that defer can be
		ensure that we never forget to unlock the mutex.

		NOTES:
		Maps are not thread-safe!
			You must lock your maps with a mutex if multiple goroutines will access it.
*/
func protectedFunction(mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock() // with defer this will be called at the end of the execution
	// do something thread-safely
}

/*
	RW Mutext
		sync.RWMutex adds two additional methods:
			- .RLock()
			- .RUnlock()
		Allows multiple readers using .RLock(), but only one can still hold a .Lock().

		This helps with performance if we have a read-intensive process.
*/
func rlockProtected(mutex *sync.RWMutex) {
	mutex.RLock()
	defer mutex.RUnlock()
	// do something thread-safely
}

func main() {
	// Mutex instantiation
	mutex := sync.Mutex{}

	protectedFunction(&mutex)

	// RWMutex instantiation
	rwMutex := sync.RWMutex{}
	rlockProtected(&rwMutex)
}
