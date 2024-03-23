package ep1

import (
	"fmt"
	"time"
)

func EP1() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}

	for _, ninja := range evilNinjas {
		go attack(ninja)
	}

	time.Sleep(time.Second * 2)
}

func attack(target string) {
	fmt.Println("Throwing ninja starts at", target)
	time.Sleep(time.Second)
}
