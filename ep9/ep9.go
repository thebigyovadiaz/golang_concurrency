package ep9

import (
	"fmt"
	"sync"
)

func EP9() {
	var numMemPieces int

	memPool := &sync.Pool{
		New: func() interface{} {
			numMemPieces++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	const numWorkers = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			mem := memPool.Get().(*[]byte)
			fmt.Sprintln("Takin some time on the resources")
			memPool.Put(mem)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("%d numMemPieces were created", numMemPieces)
}
