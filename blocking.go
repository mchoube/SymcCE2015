package main

import "some_lib"

func main() {
	// do some work

	ch := make(chan bool)
	go func(ch chan bool) {
		some_lib.blockingCall()
		ch <- true
	}(ch)

	// do some other work

	// wait for the blocking call
	<-ch
}
