package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"pack.ag/amqp"
)

func main() {
	// Create client
	client, err := amqp.Dial("amqp://127.0.0.1")
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer client.Close()

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	ctx := context.Background()

	// Send a message
	{
		// Create a sender
		sender, err := session.NewSender(
			amqp.LinkTargetAddress("/queue-name"),
		)
		if err != nil {
			log.Fatal("Creating sender link:", err)
		}

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

		// Send message
		for i := 0; i < 1000; i++ {
			err = sender.Send(ctx, amqp.NewMessage([]byte(fmt.Sprintf("Hello %d!", i+1))))
			if err != nil {
				log.Fatal("Sending message:", err)
			}
		}

		sender.Close(ctx)
		cancel()
	}
}
