package ep6

import (
	"fmt"
	"sync"
)

func EP6() {
	var beeper sync.WaitGroup
	evilNinjas := []string{"Tommy", "Johnny", "Bobby"}
	beeper.Add(len(evilNinjas))

	for _, ninja := range evilNinjas {
		go attack(ninja, &beeper)
	}

	beeper.Wait()
	fmt.Println("Mission completed")
}

func attack(ninja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja: ", ninja)
	beeper.Done()
}
