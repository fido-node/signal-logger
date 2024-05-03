package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handler(signal os.Signal) {
	if signal == syscall.SIGTERM {
		fmt.Println("Got term signal. ")
	} else if signal == syscall.SIGINT {
		fmt.Println("Got CTRL+C signal")
	} else {
		fmt.Println("Ignoring signal: ", signal)
	}
}

func main() {
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)
	exitchnl := make(chan int)

	go func() {
		for {
			s := <-sigchnl
			handler(s)
		}
	}()

	exitcode := <-exitchnl
	os.Exit(exitcode)
}
