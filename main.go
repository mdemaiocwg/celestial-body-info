package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mdemaiocwg/celestial-body-info/planet/planetpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Planet Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := planetpb.NewPlanetServiceClient(cc)

	// read Planet
	fmt.Println("Reading the planet")

	res, err2 := c.ReadPlanet(context.Background(), &planetpb.ReadPlanetRequest{PlanetId: "5f47c1ad093cc8a6577e0918"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v \n", err2)
	}

	fmt.Printf("Planet was read: %v \n", res)

}
