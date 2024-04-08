package main

import (
	"fmt"
	"time"
)

func main() {
	// Simulating log messages
	logMessages := []string{
		"Starting Selenium stanalone mode...",
		"Logs are running via the hyperexecute stanalone mode",
		"side the loop, it iterates over the logMessages slice and prints each message.",
		"After printing all messages, it sleeps for 2 seconds before printing the next set of messages. ",
		"This delay simulates the interval between fetching new log messages. You can adjust the sleep duration as needed.",
	}

	// Infinite loop to continuously print log messages
	for {
		// Loop over the log messages and print them
		for _, msg := range logMessages {
			fmt.Println(msg)
		}

		// Sleep for a short duration before printing the next set of log messages
		time.Sleep(2 * time.Second)
	}
}
