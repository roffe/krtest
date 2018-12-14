package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/roffe/krtest/pkg/server"
)

func main() {
	quitChan := make(chan bool)
	defer close(quitChan)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	srv := server.New()
	go srv.Run(ctx, quitChan)

	sig := <-sigc
	log.Printf("Got %s signal, gracefully shutting down\n", sig)
	quitChan <- true
	select {
	case <-time.After(30 * time.Second):
		return
	case sig := <-sigc:
		log.Printf("Got %s signal, forcefully quitting\n", sig)
		return
	}

}
