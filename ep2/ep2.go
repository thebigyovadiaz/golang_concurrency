package ep2

import (
	"fmt"
	"time"
)

func EP2() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	evilNinja := "Tommy"
	go attack(evilNinja)
	time.Sleep(time.Second * 2)
}

func attack(target string) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja starts at Tommy", target)
}
