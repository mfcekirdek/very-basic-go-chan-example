package main

import (
	"log"
)

func main() {
	stringChan := make(chan string)
	quitChan := make(chan bool)
	done := make(chan bool)

	go func() {
		for {
			select {
			case message := <-stringChan:
				log.Printf("Got a message from stringChan: %s\n", message)
			case <-quitChan: // If we receive anything on quitchan, we send true to the done channel and return.
				done <- true
				return
			}
		}
	}()

	log.Println("Adding some messages...")
	stringChan <- "Hello there."
	stringChan <- "Another message."
	stringChan <- "Yet another message."
	quitChan <- true

	// In the main goroutine, we’re blocked before exiting by waiting on a read from the done channel
	<-done
	log.Println("Exiting..")

}
