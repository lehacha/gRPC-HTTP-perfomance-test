package main

import (
	"fmt"
	"log"
	"time"
	"workspace"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := workspace.NewTextProcessorClient(conn)

	var timeToMake1000req []float64

	for j := 0; j < 100; j++ {
		// experiment with 1000 requests

		// experiment start time
		start := time.Now()
		for i := 0; i < 1000; i++ {
			_, err := c.Process(context.Background(), &workspace.ProcessRequest{
				Text:     "test text",
				Username: "test user name",
			})
			if err != nil {
				log.Fatalf("Error when calling Process: %s", err)
			}
		}

		// difference: now - start
		dif := time.Now().Sub(start)
		fmt.Printf("gRPC: %v \n", dif)

		timeToMake1000req = append(timeToMake1000req, float64(dif.Milliseconds()))

	}

	var sum float64
	for _, t := range timeToMake1000req {
		sum += t
	}
	fmt.Printf("[gRPC] Result: %vms\n", sum/float64(len(timeToMake1000req)))
}
