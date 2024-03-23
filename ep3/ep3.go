package ep3

import "fmt"

func EP3() {
	channel := make(chan string, 2)
	channel <- "First message"
	channel <- "Second message"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
