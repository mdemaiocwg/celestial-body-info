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

	resRead, err := c.ReadPlanet(context.Background(), &planetpb.ReadPlanetRequest{PlanetId: "5f47c1ad093cc8a6577e0918"})
	if err != nil {
		fmt.Printf("Error happened while reading: %v \n", err)
	}

	fmt.Printf("Planet was read: %v \n", resRead)

	// list Planet
	fmt.Println("Listing the planets")

	resList, err := c.ListPlanet(context.Background(), &planetpb.ListPlanetRequest{})
	if err != nil {
		fmt.Printf("Error happened while listing: %v \n", err)
	}

	fmt.Printf("Planets were listed: %v \n", resList)
}
