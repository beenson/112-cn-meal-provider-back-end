package internal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func WaitUntilShutdownSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM,
	)

	<-c
	log.Println("Shutting down...")
}
