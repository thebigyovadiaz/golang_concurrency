package ep4

import (
	"fmt"
	"math/rand"
	"time"
)

func EP4() {
	// var total int
	// channel := make(chan map[string]int)

	channel := make(chan string)
	go throwingNinjaStar(channel)

	// Type of extract channel data
	/*for message := range channel {
		fmt.Println(message)
	}*/

	// Another type of extract channel data
	for {
		message, open := <-channel
		if !open {
			break
		}

		fmt.Println(message)
	}
}

func throwingNinjaStar(channel chan string) {
	numRounds := 3

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("Your score:", score)
	}

	close(channel)
}

func throwingNinjaStarII(channel chan map[string]int, numRounds int) {
	accumulator := make(map[string]int)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		c := fmt.Sprint("round-", i)
		accumulator[c] = score
	}

	channel <- accumulator
}
