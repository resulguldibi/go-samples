// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package main

import "fmt"
import "time"

func main() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 3)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	i := 0
	for i < 2 {
		go getMessage(messages)
	}

	time.Sleep(time.Millisecond * 100)
}

func getMessage(mess chan string) {

	if message, err := <-mess; err == nil {
		fmt.Println(message)
	}

}
