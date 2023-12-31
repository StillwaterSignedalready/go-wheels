package yes

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Chatter() {
	sigChan := make(chan os.Signal, 1)
	stopTiming := time.Now().Add(time.Second * 3)

	signal.Notify(sigChan, syscall.SIGPIPE, syscall.SIGINT, os.Interrupt)
	go func() {
		<-sigChan
		os.Exit(0) // no print
	}()

	for time.Now().Before(stopTiming) { // while()
		os.Stdout.Write([]byte("y\n"))
		// os.Stdout.Sync() // flush buffer content to terminal
	}

}
