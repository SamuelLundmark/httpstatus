package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/SamuelLundmark/httpstatus/httpserver"
)

func main() {

	// Collect all arguments
	args := os.Args[1:]

	// Make sure all arguments are integers
	codes := make([]int, len(args))
	for i, arg := range args {
		v, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("argument %v: \"%s\" is not an integer\n", i, arg)
			os.Exit(1)
		}
		codes[i] = v
	}

	// Create a error channel so the interrupts can be properly handled
	errorChannel := make(chan error)

	// Setup HTTP server
	go func() {
		server := httpserver.New(":8080", codes)

		// Start HTTP server
		errorChannel <- server.ListenAndServe()
	}()

	// Capture interrupts
	go func() {
		c := make(chan os.Signal, 10)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("got signal: %s", <-c)
	}()

	// Terminate with error if error is present
	if err := <-errorChannel; err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
