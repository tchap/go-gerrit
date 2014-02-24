package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/tchap/go-gerrit/gerrit"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("Error: ")

	session, err := gerrit.Dial(nil)
	if err != nil {
		log.Fatal(err)
	}

	stream, err := session.NewEventStream()
	if err != nil {
		session.Close()
		log.Fatal(err)
	}

	signalCh := make(os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	for {
		select {
		case event := <-stream.Chan:
			fmt.Printf("%#v\n", event)

		case <-signalCh:
			stream.Close()
			session.Close()
			return
		}
	}
}
