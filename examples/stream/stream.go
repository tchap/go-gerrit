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

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	defer func() {
		if err := stream.Close(); err != nil {
			log.Print(err)
		}
		if err := session.Close(); err != nil {
			log.Print(err)
		}
	}()

	for {
		select {
		case event, ok := <-stream.Chan():
			if !ok {
				return
			}
			fmt.Printf("%#v\n", event)

		case <-signalCh:
			return
		}
	}
}
