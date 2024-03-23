package ep7

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock   sync.Mutex
	rwLock sync.RWMutex
	count  int
	iter   time.Duration
)

func EP7() {
	// basics()
	readAndWrite()
}

func basics() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Count resulted: ", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}

func readAndWrite() {
	go read()
	go write(100000)
	go read()

	time.Sleep(iter + 3*time.Second)
	fmt.Println("Done")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Reading locking")
	fmt.Println("Count resulted: ", count)
	time.Sleep(1 * time.Second)
	fmt.Println("Reading unlocking")
}

func iterUp(iteration int) int {
	return iteration % 2
}

func write(iteration int) {
	rwLock.Lock()
	defer rwLock.Unlock()

	if iteration >= 8000 {
		iter = time.Duration(iterUp(iteration))
	} else if iteration >= 5000 {
		iter = 3
	} else if iteration >= 3000 {
		iter = 2
	} else {
		iter = 1
	}

	fmt.Println("Writing locking")

	for i := 0; i < iteration; i++ {
		count += 1
	}

	time.Sleep(iter * time.Second)
	fmt.Println("Writing unlocking")
}
